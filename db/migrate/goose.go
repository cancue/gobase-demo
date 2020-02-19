package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	"gobase_demo/config"
)

func main() {
	_, yaml := config.GetStageYAML()
	dbDriver, dbString := getDbString(yaml)

	db, err := sql.Open(dbDriver, dbString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	args := os.Args[1:]
	if err := goose.Run(args[0], db, "./db/migrate/migrations", args[1:]...); err != nil {
		panic(err)
	}
}

func getDbString(yaml map[string]interface{}) (dbDriver string, dbString string) {
	db := yaml["db"].(map[string]interface{})
	connection := db["connection"].(map[string]interface{})

	dbDriver = db["driver"].(string)
	dbString = fmt.Sprintf(
		"port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		connection["port"],
		connection["host"],
		connection["user"],
		connection["password"],
		connection["database"],
	)

	return
}
