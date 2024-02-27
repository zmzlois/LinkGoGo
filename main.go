package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/joho/godotenv"
	"github.com/zmzlois/LinkGoGo/auth"
	"github.com/zmzlois/LinkGoGo/web/pages"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	fmt.Printf("Port is trying to listening on %s\n\n", port)

	app := fiber.New()

	req := &fiber.Request{}
	res := &fiber.Response{}

	app.Static("/", "./web/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Home page")
		return Render(c, pages.HomePage())
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		fmt.Println("Sign in page")
		return Render(c, pages.SignInPage())
	})

	// everything below needs to be authorized

	var cred *auth.Client = &auth.Client{
		ClientId:     os.Getenv("DISCORD_CLIENT_ID"),
		ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
		RedirectUri:  os.Getenv("DISCORD_REDIRECT_URI"),
	}

	var discordClient *auth.Client = auth.Initialise(&auth.Client{
		ClientId:     cred.ClientId,
		ClientSecret: cred.ClientSecret,
		RedirectUri:  cred.RedirectUri,
		Scopes:       []string{auth.ScopeIdentify},
	})

	app.Get("/discord-oauth", func(c *fiber.Ctx) error {
		fmt.Println("Redirecting to Discord")
		discordClient.RedirectHandler(c, res, req, "")
		fmt.Println("Status and error: ", c.Response().StatusCode())
		return nil
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
