package models

import "github.com/lucianorbr/API_Golang/db"

func Get(id int64) (todo Todo, err error) {

	// Open database connection
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Prepare statement for inserting data
	row := conn.QueryRow(`SELECT * FROM todos WHERE id = $1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
