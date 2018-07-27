package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MohitArora1/gophercises/CLI/task/cmd"
	"github.com/MohitArora1/gophercises/CLI/task/repository"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	repository.Init(dbPath)
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
