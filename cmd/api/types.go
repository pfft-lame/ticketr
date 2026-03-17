package main

import (
	"ticketr/internal/db/queries"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	cfg      cfg
	validate *validator.Validate
	db       *pgxpool.Pool
	queries  queries.Querier
	// logger
}

type cfg struct {
	port string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
