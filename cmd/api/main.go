package main

import (
	"arcade/internal/database"
	"arcade/internal/server"
	"arcade/internal/utils"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

    token, _ := utils.ReadEnvVar("TOKEN")
    db_url, err := utils.ReadEnvVar("TURSO")

    if err != nil {
        log.Println(err)
        db_url = "file:testing.db"
    }

    
    database.DB = sqlx.MustConnect("libsql", db_url + token)
	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
