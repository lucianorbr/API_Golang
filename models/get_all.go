package models

import "github.com/lucianorbr/API_Golang/db"

func GetAll() (todos []Todo, err error) {

	// Open database connection
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Prepare statement for inserting data
	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return
}
