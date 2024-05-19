package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	// "net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	user      = "shen"
	password  = ""
	port      = "localhost:5432"
	dbname    = "postgres"
	parameter = "sslmode=disable"
)

func main() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?%s", user, password, port, dbname, parameter)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()

	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected!")
}
