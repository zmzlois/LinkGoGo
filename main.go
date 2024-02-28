package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/fiber/v2/utils"
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

	sentryDSN := os.Getenv("SENTRY_DSN")

	err = sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if c, ok := hint.Context.Value(sentry.RequestContextKey).(*fiber.Ctx); ok {
					// You have access to the original Context if it panicked
					fmt.Println(utils.CopyString(c.Hostname()))
				}
			}
			fmt.Println(event)
			return event
		},
		Debug:            true,
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	app := fiber.New()

	app.Use(fibersentry.New(fibersentry.Config{
		Repanic:         true,
		WaitForDelivery: true,
	}))

	// req := &fiber.Request{}
	res := &fiber.Response{}
	resWrite := res.BodyWriter()

	app.Static("/", "./web/assets")

	var cred *auth.Client = &auth.Client{
		ClientId:     os.Getenv("DISCORD_CLIENT_ID"),
		ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
		RedirectUri:  os.Getenv("DISCORD_REDIRECT_URI"),
	}

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Home page")
		return Render(c, pages.HomePage())
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		fmt.Println("They hit sign in page!")
		return Render(c, pages.SignInPage())
	})

	// everything below needs to be authorized

	// handle discord sign in

	var discordClient *auth.Client = auth.Initialise(&auth.Client{
		ClientId:     cred.ClientId,
		ClientSecret: cred.ClientSecret,
		RedirectUri:  cred.RedirectUri,
		Scopes:       []string{auth.ScopeIdentify},
	})

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/discord-oauth":    discordClient.AuthUrl,
			"/discord-redirect": "/edit",
		},
		StatusCode: 307,
	}))
	app.Get("/discord-oauth", func(c *fiber.Ctx) error {
		fmt.Println("Redirecting to discord sign in")
		return nil
	})

	app.Get("/discord-redirect", func(c *fiber.Ctx) error {
		fmt.Println("After log in redirected from discord")

		var (
			// get the code from the query
			code = c.Query("code")

			accessToken, _ = discordClient.GetOnlyAccessToken(code)

			userData, _ = auth.GetDiscordUserData(accessToken)
		)
		fmt.Println("Code: ", code)

		fmt.Fprint(resWrite, userData)
		return Render(c, pages.EditPage())
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
