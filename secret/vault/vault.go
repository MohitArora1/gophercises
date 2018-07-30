package vault

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/MohitArora1/gophercises/secret/cipher"
)

// Vault is struct used to store API's keys
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// GetVault is used to get vault struct to perform get set operation
func GetVault(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return err
	}
	defer f.Close()
	return v.reader(f)
}
func (v *Vault) reader(f io.Reader) error {
	dec := json.NewDecoder(f)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) save(key, value string) error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	hex, err := cipher.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}
	v.keyValues[key] = hex
	return v.writer(f)
}
func (v *Vault) writer(f io.Writer) error {
	enc := json.NewEncoder(f)
	return enc.Encode(v.keyValues)
}

// Set function is used to set the key value in secret file
// it takes key and value as a string and return error
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.load()
	err := v.save(key, value)
	if err != nil {
		return errors.New("Not able to save into file")
	}
	return nil
}

// Get function is used to get the value from secret file
// it takes key as a value and return plaintext as string and error
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.load()
	if err != nil {
		return "", errors.New("File not found")
	}
	hex, ok := v.keyValues[key]
	if ok != true {
		return "", errors.New("Key not found")
	}
	value, err := cipher.Decrypt(v.encodingKey, hex)
	if err != nil {
		return "", errors.New("Not able to decrypt")
	}
	return value, nil
}
