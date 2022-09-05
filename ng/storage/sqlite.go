package storage

import (
	"database/sql"
	"errors"
	"os"

	"github.com/leo-alvarenga/to-go/shared/task"
	_ "github.com/mattn/go-sqlite3"
)

const dbName = "./database.db"

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", dbName)

	if err != nil {
		os.Create(dbName)

		return nil
	}

	err = db.Ping()

	return db
}

func createTaskTable(db *sql.DB) {
	w, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY,
			title TEXT,
			description TEXT,
			priority TEXT,
			status TEXT,
			createdIn TEXT,
			finishedIn TEXT
		)
	`)

	w.Exec()
}

func readFromDB(db *sql.DB, taskSlice *[]task.Task, priority string, done chan bool) {
	createTaskTable(db)
	var rows *sql.Rows

	if priority == "" {
		rows, _ = db.Query(`SELECT * FROM tasks`)
	} else {
		rows, _ = db.Query(`SELECT * FROM tasks WHERE priority = ?`, priority)
	}

	var t task.Task
	var id int
	for rows.Next() {
		rows.Scan(
			&id,
			&t.Title,
			&t.Description,
			&t.Priority,
			&t.Status,
			&t.CreatedIn,
			&t.FinishedIn,
		)

		*taskSlice = append(*taskSlice, t)
	}

	done <- true
}

func RetriveTasksFromSQLite(h, m, l *[]task.Task) error {
	db := ConnectDB()

	if db == nil {
		return errors.New("Unable to connect to database. Please, try again.")
	}

	defer db.Close()

	low, med, high := make(chan bool), make(chan bool), make(chan bool)

	go readFromDB(db, h, "high", high)
	go readFromDB(db, m, "medium", med)
	go readFromDB(db, l, "low", low)

	ok := (<-high && <-med && <-low)

	if !ok {
		return errors.New("One or more priorities has 0 tasks assigned.")
	}

	return nil
}

func WriteToSQLite(t task.Task) error {
	db := ConnectDB()

	if db == nil {
		return errors.New("Unable to connect to database. Please, try again.")
	}

	defer db.Close()

	s, err :=
		db.Prepare(`INSERT INTO tasks (title, description, priority, status, createdIn, finishedIn) VALUES (?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return errors.New("Unable to insert task. Please, try again.")
	}

	_, err = s.Exec(
		t.Title,
		t.Description,
		t.Priority,
		t.Status,
		t.CreatedIn,
		t.FinishedIn)

	if err != nil {
		return errors.New("Unable to finish insertion. Please, try again.")
	}

	return nil
}
