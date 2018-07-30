package cmd

import (
	"testing"
)

func TestSet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		key         string
		plainText   string
		expected    string
	}{
		{encodingKey: "123", key: "twitter", plainText: "hello", expected: `saved key success
		`},
	}
	for _, test := range testSuit {
		encodingKey = test.encodingKey
		args := []string{
			test.key,
			test.plainText,
		}
		setCmd.Run(setCmd, args)
	}
}

func TestGet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		key         string
		plainText   string
	}{
		{encodingKey: "123", key: "twitter", plainText: "hello"},
		{encodingKey: "123", key: "google", plainText: "hello"},
	}
	for _, test := range testSuit {
		encodingKey = test.encodingKey
		args := []string{
			test.key,
		}
		getCmd.Run(getCmd, args)
	}
}

func TestRoot(t *testing.T) {
	err := Execute()
	if err != nil {
		t.Error("error in root")
	}
}
