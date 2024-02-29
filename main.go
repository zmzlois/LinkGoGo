package main

// TODO: middleware if you don't have any cookies "github or discord", bounce back to login page
// TODO: how are you getting the cookie

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
	"github.com/zmzlois/LinkGoGo/handlers"
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
	// res := &fiber.Response{}
	// resWrite := res.BodyWriter()

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
		fmt.Println("They hit log in page!")

		code := c.Cookies("discord_code")

		if len(code) > 1 {
			return c.Redirect("/user/edit")
		}
		return Render(c, pages.LogInPage())
	})

	// everything below needs to be authorized

	// handle discord sign in

	var dc *auth.Client = auth.DiscordInitialise(&auth.Client{
		ClientId:     cred.ClientId,
		ClientSecret: cred.ClientSecret,
		RedirectUri:  cred.RedirectUri,
		Scopes:       []string{auth.ScopeIdentify},
	})

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/discord-oauth": dc.AuthUrl,
		},
		StatusCode: fiber.StatusTemporaryRedirect,
	}))

	app.Get("/discord-oauth", func(c *fiber.Ctx) error {
		fmt.Println("Redirecting to discord sign in")
		return nil
	})

	app.Get("/discord-redirect", func(c *fiber.Ctx) error {
		fmt.Println("After log in redirected from discord")

		var (
			// get the code from the query
			code = c.Queries()["code"]

			accessToken = dc.GetAccessToken(c)

			// userData, _ = auth.GetDiscordUserData(accessToken)
		)

		fmt.Println("Code: ", code)
		// cookie := new(fiber.Cookie)

		cookie := fiber.Cookie{
			Name:    "discord_code",
			Value:   code,
			Expires: time.Now().Add(120 * time.Hour),
		}

		c.Cookie(&cookie)

		fmt.Println("Access Token: ", accessToken)

		// fmt.Println(resWrite, "UserData", userData)

		// verify
		return c.Redirect("/user/edit")
	})

	user := app.Group("/user", func(c *fiber.Ctx) error {

		fmt.Println("Implementing middleware")
		handlers.ProtectedRoutes(c)
		fmt.Println("Protected routes passed!")
		return c.Next()
	})

	user.Get("/edit", func(c *fiber.Ctx) error {

		code := c.Cookies("discord_code")
		fmt.Println("Code: ", code)
		return Render(c, pages.EditPage())
	})

	user.Get("/account", func(c *fiber.Ctx) error {

		return nil
	})

	app.Post("/logout", func(c *fiber.Ctx) error {

		fmt.Println("We hit log out route!")

		cookie := fiber.Cookie{
			Name:  "discord_code",
			Value: "",
		}

		fmt.Println("Cookie: ")
		c.Cookie(&cookie)

		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	})

	log.Fatal(app.Listen(port))

	fmt.Println("Listening on 3000")
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
