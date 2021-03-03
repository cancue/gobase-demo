package router

import (
	"github.com/cancue/gobase/controller"
	"github.com/gofiber/fiber/v2"
	"gobase_demo/interface/controller/static"
)

// Router links path to controllers before the server is started.
func Router(app *fiber.App) {
	staticAPI(app)
}

func staticAPI(app *fiber.App) {
	app.Get("/status", static.Status)
	app.Get("/string", static.String)
	app.Get("/json", controller.RouteJSON(new(static.JSON)))
}
