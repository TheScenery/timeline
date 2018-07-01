package main

import (
	"fmt"
	"os"

	"github.com/TheScenery/timeline/server/database"
	"github.com/jackc/pgx"
)

func main() {

	db, err := database.InitDatabase()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	connPool := db.GetConnPool()

	result, err := connPool.Exec(`CREATE TABLE events (
		subject text,
		description text
	);`)

	if err != nil && err.(pgx.PgError).Code != "42P07" {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(result)

	result, err = connPool.Exec("insert into events(subject, description) values($1, $2)", "Test Subject", "This is a test subject")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(result)

	rows, err := connPool.Query("select * from events")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var subject string
		var description string
		err := rows.Scan(&subject, &description)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s %s\n", subject, description)
	}

	fmt.Println(rows.FieldDescriptions())
}
