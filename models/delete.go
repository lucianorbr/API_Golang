package models

import "github.com/lucianorbr/API_Golang/db"

func Delete(id int64) (int64, error) {

	// Open database connection
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// Prepare statement for inserting data
	res, err := conn.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
