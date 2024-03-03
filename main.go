package main

// TODO: middleware if you don't have any cookies "github or discord", bounce back to login page
// TODO: how are you getting the cookie

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	//dsc "github.com/realTristan/disgoauth"

	"github.com/zmzlois/LinkGoGo/auth"
	dsc "github.com/zmzlois/LinkGoGo/auth"
	"github.com/zmzlois/LinkGoGo/handlers"
	"github.com/zmzlois/LinkGoGo/monitor"
	m "github.com/zmzlois/LinkGoGo/monitor"
	"github.com/zmzlois/LinkGoGo/web/pages"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	fmt.Printf("Port is trying to listening on %s\n\n", port)

	// Initialize sentry
	monitor.SentryInit()
	// db := handlers.DatabaseClient()

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	app := chi.NewRouter()

	app.Use(middleware.Logger)

	// Serving static files: Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()

	filesDir := http.Dir(filepath.Join(workDir, "/web/assets"))
	FileServer(app, "/", filesDir)

	app.Get("/", m.Sh.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		pages.HomePage().Render(r.Context(), w)
	}))

	app.Get("/login", m.Sh.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		pages.LogInPage().Render(r.Context(), w)
	}))

	var cred = auth.AuthCred()

	var ds *dsc.Client = dsc.Init(&dsc.Client{
		RedirectURI: cred.RedirectURI,
		// RefreshRedirectURI: cred.RefreshRedirectURI,
		ClientID:     cred.ClientID,
		ClientSecret: cred.ClientSecret,
		Scopes:       []string{auth.ScopeIdentify},
	})

	//var state string

	// Once user hit discord login button
	app.Get("/discord-oauth", func(w http.ResponseWriter, r *http.Request) {
		ds.RedirectHandler(w, r, "")
	})

	// On discord's authentication page they will be redirected to this page after authentication

	app.Get("/discord-redirect", func(w http.ResponseWriter, r *http.Request) {
		handlers.AuthenticationHandler(ds)(w, r)
		// fmt.Printf("Data: %s\n Access Token: %s\n Refresh Token: %s\n Expire In: %d\n", data, accessToken, refreshToken, expireIn)

	})

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

	log.Fatal(http.ListenAndServe(port, app))

	fmt.Println("Listening on 3000")
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {

	fmt.Printf("Path: %s | Root: %s\n", path, root)
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
