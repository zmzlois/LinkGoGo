package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/database"
	"github.com/zmzlois/LinkGoGo/pkg/handler"
	"github.com/zmzlois/LinkGoGo/pkg/monitor"
	"github.com/zmzlois/LinkGoGo/pkg/service"
	"github.com/zmzlois/LinkGoGo/pkg/utils"
	"github.com/zmzlois/LinkGoGo/web/pages"
)

func SetupRouter(app chi.Router) {
	db := database.Connection()

	ds := auth.DiscordInit()

	state, err := utils.StateGenerator()
	if err != nil {
		panic("Error generating state")
	}

	discordOAuth := service.NewAuthService(db, ds, state)
	discordOAuthService := handler.NewAuthHandler(*discordOAuth)

	// discordOAuthRedirectService := service.NewAuthService(db, ds, state)

	indexHandler := monitor.SentryHandler.HandleFunc(handler.Index)
	loginHandler := monitor.SentryHandler.HandleFunc(handler.Login)

	discordOAuthHandler := monitor.SentryHandler.HandleFunc(handler.OAuthHandler)
	discordOAuthRedirectHandler := monitor.SentryHandler.HandleFunc(discordOAuthService.OAuthRedirectHandler)
	// discordOAuthHandler := monitor.SentryHandler.HandleFunc(handler.NewDiscordOAuthHandler(*discordOAuthService))
	// discordOAuthRedirectHandler := monitor.SentryHandler.HandleFunc(handler.NewDiscordOAuthRedirect(*discordOAuthRedirectService))

	// Index route
	app.Get("/", handler.Index)

	// Login route
	app.Get("/login", monitor.SentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		pages.LogInPage().Render(r.Context(), w)
	}))
	// Auth route
	app.Post("/discord-oauth", handler.OAuthHandler)

	app.Get("/", indexHandler)
	app.Get("/login", loginHandler)
	// var state string

	// Once user hit discord login button
	app.Get("/discord-oauth", discordOAuthHandler)

	// On discord's authentication page they will be redirected to this page after authentication
	// app.Get("/discord-redirect", func(w http.ResponseWriter, r *http.Request) {
	// hdl.AuthenticationHandler(ds)(w, r)
	// })

	app.Get("/discord-redirect", discordOAuthRedirectHandler)

	app.Group(func(r chi.Router) {
		// r.Use(handlers.ProtectedRoutes(r))

		r.Get("/edit", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("URL: %s\n", r.URL.Path)

			pages.EditPage().Render(r.Context(), w)
		})

		r.Get("/account", func(w http.ResponseWriter, r *http.Request) {
		})

		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusFound)
		})
	})

	app.Get("/failed", func(w http.ResponseWriter, r *http.Request) {
		// pages.FailedPage().Render(r.Context(), w)

		// set time out and redirect to home page
		time.Sleep(5 * time.Second)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	app.NotFound(func(w http.ResponseWriter, r *http.Request) {
		pages.NotFound().Render(r.Context(), w)
	},
	)

	app.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		pages.NotFound().Render(r.Context(), w)
	},
	)
}
