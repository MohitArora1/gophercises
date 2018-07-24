// this package deals with all the repository based operations
package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// this function used to initialized the db
// as well create table if not created.
func Init(dbPath string) {
	var err error

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec("create table if not exists tasks(task text primary key)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// this function used to insert the task into db
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

// this function return task from db which is not done
// return list to string and error
func ReadNotCompletedTaskFromDB() ([]string, error) {
	var tasks []string
	var task string
	stmt, err := db.Prepare("select task from tasks")
	if err != nil {
		return tasks, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return tasks, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&task)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// this function remove the task from db
// return type status and error
func MarkTaskAsDone(ids []int) (bool, error) {
	tasks, err := ReadNotCompletedTaskFromDB()
	if err != nil {
		return false, err
	}
	var deleteTask []string
	for _, id := range ids {
		deleteTask = append(deleteTask, tasks[id-1])
	}
	for _, task := range deleteTask {
		stmt, err := db.Prepare("delete from tasks where task=?")
		if err != nil {
			return false, err
		}
		_, err = stmt.Exec(task)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
