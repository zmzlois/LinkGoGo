package main

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/zmzlois/go-tooo/internal/templates/pages"
)

func main() {
	port := ":3000"

	fmt.Printf("Port is trying to listening on %s\n\n", port)

	app := fiber.New()

	app.Static("/", "./internal/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, pages.HomePage())
	})

	app.Listen(port)

	fmt.Println("Listening on 3000")
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
