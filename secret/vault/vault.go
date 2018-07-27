package vault

import (
	"os"
	"sync"
)

type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

func GetVault(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) load() {

}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	return nil
}

func (v *Vault) Get(key string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	return nil
}
