package main

import (
	"github.com/cancue/gobase/config"
	"github.com/cancue/gobase/db/pg"
)

func main() {
	conf := config.Set("..")
	pg.Goose(conf)
}
