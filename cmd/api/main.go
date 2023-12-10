package main

import (
	"arcade/internal/server"
    "arcade/internal/database"
    "github.com/jmoiron/sqlx"
    
    _ "github.com/mattn/go-sqlite3"
)

func main() {

    database.DB = sqlx.MustConnect("sqlite3", "testing.db")
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
