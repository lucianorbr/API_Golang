package models

import "github.com/lucianorbr/API_Golang/db"

func Insert(todo Todo) (id int64, err error) {

	// Open database connection
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// Prepare statement for inserting data
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
