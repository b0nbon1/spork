package main

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Views interface {
    Load() error
    Render(io.Writer, string, interface{}, ...string) error
}

type Count struct {
	Count int
}

func main() {
	// Create a new engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
        Views: engine,
    })

	count := Count{Count: 0}

	app.Get("/", func(c *fiber.Ctx) error {
		count.Count++
		return c.Render("index", count)
	})

	app.Listen(":6942")
}
