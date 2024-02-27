package main

import (
	"database/sql"
	"log"
	"simplebank/api"
	db "simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:4233/simple_bank?sslmode=disable"
	serveraddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("DB connection failed", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serveraddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
