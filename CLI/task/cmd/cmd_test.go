package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
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
		args []string
	}{
		{args: []string{"go", "to", "gym"}},
		{args: []string{"go", "to", "office"}},
		{args: []string{"go", "to", "office"}},
		{args: []string{"clean", "dishes"}},
	}
	for _, test := range testSuit {
		addCmd.Run(addCmd, test.args)
	}
	expected := `Added "go to gym" 
Added "go to office" 
Not able to add task to todo
Added "clean dishes" 
`
	b, _ := ioutil.ReadFile(file.Name())
	output := string(b)
	if expected != output {
		t.Error("error in add command")
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
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file

	initial()
	testSuit := []struct {
		args []string
	}{
		{args: []string{"1", "2", "3", "a"}},
		{args: []string{"1"}},
	}
	for _, test := range testSuit {
		doCmd.Run(doCmd, test.args)
	}
	expected := `enable to parse the id a
done tasks

[0] these ids not exist and rest mark as done
`
	b, _ := ioutil.ReadFile(file.Name())
	output := string(b)
	if expected != output {
		t.Error("error in do command")
	}
	os.Stdout = old
}

func TestRoot(t *testing.T) {
	err := Execute()
	if err != nil {
		t.Error("error in root command")
	}
}
