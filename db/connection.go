package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/lucianorbr/API_Golang/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%S sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgress", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
