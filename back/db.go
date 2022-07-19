package main

import "github.com/jackc/pgx"

var db *pgx.Conn

func initDb() (err error) {
	db, err = pgx.Connect(config.Db)
	return
}
