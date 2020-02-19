package main

import (
	"strconv"
	"time"

	"github.com/cancue/gobase"
	"github.com/labstack/echo/v4/middleware"

	"gobase_demo/config"
	"gobase_demo/router"
)

func main() {
	stage, yaml := config.GetStageYAML()

	server := gobase.NewWithConfig(newConfig(stage, yaml))
	setMiddlewares(server, yaml)
	router.Apply(server)

	server.Start()
}

func newConfig(stage string, yaml map[string]interface{}) (config *gobase.Config) {
	server := yaml["server"].(map[string]interface{})
	timeout := server["timeout"].(map[string]interface{})
	config = &gobase.Config{
		Stage:        stage,
		Name:         server["name"].(string),
		Port:         ":" + strconv.Itoa(server["port"].(int)),
		ReadTimeout:  time.Duration(timeout["read"].(int)) * time.Second,
		WriteTimeout: time.Duration(timeout["read"].(int)) * time.Second,
	}

	return
}

func setMiddlewares(server *gobase.Server, yaml map[string]interface{}) {
	allowOrigins := new([]string)
	for _, v := range yaml["allow-origins"].([]interface{}) {
		*allowOrigins = append(*allowOrigins, v.(string))
	}

	server.Use(middleware.Secure())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: *allowOrigins,
	}))
	server.Use(middleware.BodyLimit("0.1M"))
	server.Use(middleware.RequestID())
	server.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           middleware.DefaultSkipper,
		StackSize:         1 << 10, // 1 KB
		DisableStackAll:   true,
		DisablePrintStack: true,
	}))

	if server.Config.Stage == "local" {
		server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		}))
	}
}
