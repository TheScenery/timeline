package database

import (
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

type Database struct {
	connPool *pgx.ConnPool
}

func InitDatabase() (*Database, error) {
	const host string = "localhost"
	const database string = "timeline"
	port, err := strconv.ParseUint(os.Getenv("PGPORT"), 10, 16)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return &Database{connPool: connPool}, nil
}

func (database *Database) GetConnPool() *pgx.ConnPool {
	return database.connPool
}
