package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/zmzlois/LinkGoGo/pkg/auth"
	"github.com/zmzlois/LinkGoGo/pkg/database"
	"github.com/zmzlois/LinkGoGo/pkg/handler"
	"github.com/zmzlois/LinkGoGo/pkg/middleware"
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

	indexHandler := handler.Index
	loginHandler := handler.Login

	discordOAuthHandler := handler.OAuthHandler
	discordOAuthCallbackHandler := discordOAuthService.OAuthCallbackHandler

	// Index route
	app.Get("/", handler.Index)

	// Auth route
	app.Post("/discord-oauth", discordOAuthHandler)

	app.Get("/", indexHandler)
	app.Get("/login", loginHandler)

	// Once user hit discord login button
	app.Get("/discord-oauth", discordOAuthHandler)

	// On discord's authentication page they will be redirected to this page after authentication
	// app.Get("/discord-redirect", func(w http.ResponseWriter, r *http.Request) {
	// hdl.AuthenticationHandler(ds)(w, r)
	// })

	app.Get("/discord-callback", discordOAuthCallbackHandler)

	app.Get("/unauthorised", handler.UnauthorisedHandler)

	linkService := service.NewLinkService(db)

	app.Group(func(r chi.Router) {
		tr := &handler.TodosController{
			LinkService: linkService,
		}
		r.Use(middleware.AuthMiddleware)

		r.Get("/edit", tr.Edit)
		r.Get("/edit/link", tr.CreateLinkHandler)
		r.Get("/create", tr.CreateButtonHandler)

		// r.Get("/edit", monitor.SentryHandler.HandleFunc(handler.EditPageHandler))
		r.Post("/edit/{param}", handler.UrlHandler)
		r.Get("/account", func(w http.ResponseWriter, r *http.Request) {
		})
		r.Post("/create", tr.EditLinkHandler)
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/", http.StatusFound)
		})

	})

	app.Get("/logout", handler.LogoutHandler)

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
