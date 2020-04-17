package main

import (
	"github.com/cancue/gobase"
	"github.com/cancue/gobase/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gobase_demo/router"
)

func main() {
	config := config.Set(".")

	allowOrigins := new([]string)
	for _, v := range config.YAML["allow-origins"].([]interface{}) {
		*allowOrigins = append(*allowOrigins, v.(string))
	}

	middlewares := []echo.MiddlewareFunc{
		middleware.Secure(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: *allowOrigins,
		}),
		middleware.BodyLimit("0.1M"),
		middleware.RequestID(),
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			Skipper:           middleware.DefaultSkipper,
			StackSize:         1 << 10, // 1 KB
			DisableStackAll:   true,
			DisablePrintStack: false,
		}),
	}

	if config.Stage == "local" {
		middlewares = append(
			middlewares,
			middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
			}),
		)
	}

	gb := gobase.Server{
		Config:      config,
		Router:      router.Apply,
		Middlewares: middlewares,
	}

	gb.Start()
}
