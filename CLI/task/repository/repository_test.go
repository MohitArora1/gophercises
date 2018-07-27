package repository

import (
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func initial() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	Init(dbPath)
}
func TestInsertIntoDB(t *testing.T) {
	initial()
	_, err := InsertIntoDB("go to gym")
	if err != nil {
		t.Errorf("error in inserting data")
	}
}

func TestReadNotCompletedTaskFromDB(t *testing.T) {
	initial()
	_, err := ReadNotCompletedTaskFromDB()
	if err != nil {
		t.Errorf("error in reading data")
	}
}

func TestMarkTaskAsDone(t *testing.T) {
	initial()
	var ids = []int{1, 2}
	_, err := MarkTaskAsDone(ids)
	if err != nil {
		t.Error("error in doing task")
	}
}
