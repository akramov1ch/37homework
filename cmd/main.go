package main

import (
	"fmt"
	"37hw/api"
	"37hw/config"
	"37hw/storage"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := config.Load(".")

	var psqlUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.DbHost,
		cfg.Postgres.DbPort,
		cfg.Postgres.DbUser,
		cfg.Postgres.DbPassword,
		cfg.Postgres.DbName,
	)

	db, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatal("failed connecting database : ", err)
	}
	storage := storage.NewStoragePg(db)

	server := api.New(api.Option{
		Storage: storage,
	})

	if err := server.Run(":8080"); err != nil {
		log.Fatal("Failed to run HTTP server:  ", err)
		panic(err)
	}

	log.Print("Server stopped")

}
