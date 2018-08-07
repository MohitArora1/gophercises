package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestSet(t *testing.T) {
	file, err := os.Create("./test.txt")
	if err != nil {
		t.Error("Not able to open file")
	}
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	testSuit := []struct {
		encodingKey string
		key         string
		plainText   string
		expected    string
	}{
		{encodingKey: "123", key: "twitter", plainText: "hello", expected: "saved key success"},
	}
	for _, test := range testSuit {
		encodingKey = test.encodingKey
		args := []string{
			test.key,
			test.plainText,
		}
		setCmd.Run(setCmd, args)
		file.Seek(0, 0)
		b, err := ioutil.ReadAll(file)
		if err != nil {
			t.Error("error in converting into bytes")
		}
		match, err := regexp.Match(test.expected, b)
		if err != nil {
			t.Error("err in regex")
		}
		if !match {
			t.Error("error")
		}
	}

	os.Stdout = old
}

func TestGet(t *testing.T) {
	file, err := os.Create("test.txt")
	if err != nil {
		t.Error("error in creating file")
	}
	defer file.Close()
	old := os.Stdout
	os.Stdout = file
	testSuit := []struct {
		encodingKey string
		key         string
		expected    string
	}{
		{encodingKey: "123", key: "twitter", expected: "hello"},
		{encodingKey: "123", key: "google", expected: "Key not found"},
		{encodingKey: "123", key: "facebook", expected: "Key not found"},
	}
	for _, test := range testSuit {
		encodingKey = test.encodingKey
		args := []string{
			test.key,
		}
		getCmd.Run(getCmd, args)
		file.Seek(0, 0)
		b, err := ioutil.ReadAll(file)
		if err != nil {
			t.Error("error in converting into bytes")
		}
		match, err := regexp.Match(test.expected, b)
		if err != nil {
			t.Error("err in regex")
		}
		if !match {
			t.Error("error")
		}

	}
	os.Stdout = old
}

func TestRoot(t *testing.T) {
	err := Execute()
	if err != nil {
		t.Error("error in root")
	}
}
