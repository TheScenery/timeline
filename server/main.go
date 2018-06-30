package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

func main() {
	const host string = "localhost"
	const database string = "timeline"
	port, err := strconv.ParseUint(os.Getenv("PGPORT"), 10, 16)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")

	connconfig := pgx.ConnConfig{
		Host:     host,
		Port:     uint16(port),
		Database: database,
		User:     user,
		Password: password,
	}

	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: connconfig,
	}

	connPool, err := pgx.NewConnPool(connPoolConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result, err := connPool.Exec(`CREATE TABLE events (
		subject text,
		description text
	);`)

	if err != nil && err.(pgx.PgError).Code != "42P07" {
		fmt.Println(err)
		os.Exit(-1)
	}

	connPool.Exec("insert into events(subject, description) values($1, $2)", "Test Subject", "This is a test subject")

	fmt.Println(result)

	fmt.Println(connPoolConfig)
}
