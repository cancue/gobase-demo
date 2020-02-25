package db

import (
	"fmt"

	"github.com/go-pg/pg/v9"

	"gobase_demo/config"
)

var db *pg.DB

func init() {
	_, yaml := config.GetStageYAML()
	connection := yaml["db"].(map[string]interface{})["connection"].(map[string]interface{})

	db = pg.Connect(&pg.Options{
		Database: connection["database"].(string),
		Addr:     fmt.Sprintf("%s:%d", connection["host"], connection["port"]),
		User:     connection["user"].(string),
		Password: connection["password"].(string),
	})
}

// Get returns go-pg DB which is safe for concurrent use by multiple goroutines and maintains its own connection pool.
func Get() *pg.DB {
	return db
}
