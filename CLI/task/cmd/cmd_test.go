package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/MohitArora1/gophercises/CLI/task/repository"
	homedir "github.com/mitchellh/go-homedir"
)

func initial() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	repository.Init(dbPath)
}
func TestAddCLI(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	initial()
	testSuit := []struct {
		args     []string
		expected string
	}{
		{args: []string{"go", "to", "gym"}, expected: `Added "go to gym"`},
		{args: []string{"go", "to", "office"}, expected: `Added "go to office"`},
		{args: []string{"go", "to", "office"}, expected: `Not able to add task to todo`},
		{args: []string{"clean", "dishes"}, expected: `Added "clean dishes"`},
	}
	for _, test := range testSuit {
		addCmd.Run(addCmd, test.args)
		file.Seek(0, 0)
		b, _ := ioutil.ReadFile(file.Name())
		match, err := regexp.Match(test.expected, b)
		if err != nil {
			t.Error("error in regex")
		}
		if !match {
			t.Error("error")
		}
	}
	os.Stdout = old
}

func TestListCLI(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	initial()
	args := []string{}
	listCmd.Run(listCmd, args)
	expected := `1. go to gym
2. go to office
3. clean dishes
`
	b, _ := ioutil.ReadFile(file.Name())
	output := string(b)
	if expected != output {
		t.Error("error in list command")
	}
	os.Stdout = old
}

func TestDoCLI(t *testing.T) {
	file, _ := os.Create("./test.txt")
	defer file.Close()
	//defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file

	initial()
	testSuit := []struct {
		args     []string
		expected string
	}{
		{args: []string{"1", "2", "3", "a"}, expected: "enable to parse the id a"},
		{args: []string{"1"}, expected: "these ids not exist and rest mark as done"},
	}
	for _, test := range testSuit {
		doCmd.Run(doCmd, test.args)
		file.Seek(0, 0)
		b, _ := ioutil.ReadFile(file.Name())
		match, err := regexp.Match(test.expected, b)
		if err != nil {
			t.Error("error in regex")
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
		t.Error("error in root command")
	}
}
