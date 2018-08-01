// Package repository deals with all the repository based operations
package repository

import (
	"database/sql"

	// using this package as a driver for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Init function used to initialized the db
// as well create table if not created.
func Init(dbPath string) error {
	var err error

	db, err = sql.Open("sqlite3", dbPath)
	if err == nil {
		_, err = db.Exec("create table if not exists tasks(task text primary key)")
	}
	return err
}

// InsertIntoDB function used to insert the task into db
// and return status and error
func InsertIntoDB(task string) (bool, error) {
	stmt, _ := db.Prepare("insert into tasks(task) values(?)")
	defer stmt.Close()
	_, err := stmt.Exec(task)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ReadNotCompletedTaskFromDB function return task from db which is not done
// return list to string and error
func ReadNotCompletedTaskFromDB() ([]string, error) {
	var tasks []string
	var task string
	stmt, err := db.Prepare("select task from tasks")
	if err == nil {
		rows, err := stmt.Query()
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&task)
				tasks = append(tasks, task)
			}
		}
	}
	return tasks, err
}

// MarkTaskAsDone function remove the task from db
// return type status and error
func MarkTaskAsDone(ids []int) ([]int, error) {
	var notExist []int
	tasks, err := ReadNotCompletedTaskFromDB()
	if err == nil {
		var deleteTask []string
		for _, id := range ids {
			if id-1 < len(tasks) {
				deleteTask = append(deleteTask, tasks[id-1])
			} else {
				notExist = append(notExist, id-1)
			}
		}
		for _, task := range deleteTask {
			stmt, err := db.Prepare("delete from tasks where task=?")
			if err == nil {
				_, err = stmt.Exec(task)
			}
		}
	}
	return notExist, err
}
