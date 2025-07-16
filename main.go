package main

import (
	"context"
	"log"

	"github.com/aram88/simplebank/api"
	db "github.com/aram88/simplebank/db/sqlc"
	"github.com/aram88/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".s")
	if err != nil {
		log.Fatal("cannnot load configurations", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
