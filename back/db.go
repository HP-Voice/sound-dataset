package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func initDb() error {
	ctx := context.Background()
	cfg, err := pgxpool.ParseConfig(config.Db)
	if err != nil {
		return err
	}
	db, err = pgxpool.ConnectConfig(ctx, cfg)
	return err
}
