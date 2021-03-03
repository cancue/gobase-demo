package static

import (
	"github.com/gofiber/fiber/v2"
)

// Status .
func Status(ctx *fiber.Ctx) (err error) {
	ctx.SendStatus(200)
	return
}

// String .
func String(ctx *fiber.Ctx) (err error) {
	ctx.SendString("Hello, World 👋!")
	return
}

// JSON .
type JSON struct{}

func (req *JSON) Exec(ctx *fiber.Ctx) (result interface{}, err error) {
	result = map[string]string{
		"message": "Hello, World 👋!",
	}
	return
}
