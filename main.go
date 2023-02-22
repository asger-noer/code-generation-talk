package main

import (
	"database/sql"
	"fmt"

	"github.com/AsgerNoer/talks/write-less-code/internal/database"
	"github.com/AsgerNoer/talks/write-less-code/internal/httpserver"
	"github.com/rs/zerolog/log"
)

const dbDialect string = "postgres"

func main() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		"postgres",
		"postgres",
		"postgres",
		"postgres",
		"disable",
	)

	db, err := sql.Open(dbDialect, connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("opening database")
	}

	defer db.Close()

	store, err := database.New(dbDialect, db)
	if err != nil {
		log.Fatal().Err(err).Msg("establishing database connection")
	}

	server := httpserver.NewServer(store)

	if err := server.Start("0.0.0.0:3000"); err != nil {
		log.Fatal().Err(err).Msg("starting service failed")
	}

}
