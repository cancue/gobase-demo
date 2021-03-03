package main

import (
	_ "gobase_demo/config"

	"github.com/cancue/gobase"
	"github.com/cancue/gobase/config"

	"gobase_demo/interface/router"
)

func main() {
	gobase.Start(config.Get(), router.Router)
}
