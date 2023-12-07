package main

import (
	"arcade/internal/server"
    "arcade/internal/database"
    "github.com/jmoiron/sqlx"
)

func main() {

    database.DB = sqlx.MustConnect("sqlite3", "prod.db")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
