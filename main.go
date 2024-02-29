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
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/zmzlois/LinkGoGo/auth"
	"github.com/zmzlois/LinkGoGo/web/pages"
	"golang.org/x/oauth2"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	fmt.Printf("Port is trying to listening on %s\n\n", port)

	sentryDSN := os.Getenv("SENTRY_DSN")

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					fmt.Println("Request URL: %s", req.URL)
				}
			}

			return event
		},
		Debug:            true,
		AttachStacktrace: true,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic:         true,
		WaitForDelivery: true,
		Timeout:         2 * time.Second,
	})

	app := chi.NewRouter()

	app.Use(middleware.Logger)

	var cred *auth.Client = &auth.Client{
		ClientId:     os.Getenv("DISCORD_CLIENT_ID"),
		ClientSecret: os.Getenv("DISCORD_CLIENT_SECRET"),
		RedirectUri:  os.Getenv("DISCORD_REDIRECT_URI"),
	}

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()

	fmt.Println("Work directory: ", workDir)
	filesDir := http.Dir(filepath.Join(workDir, "/web/assets"))
	fmt.Println("Root: ", filesDir)
	FileServer(app, "/", filesDir)

	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.HomePage().Render(r.Context(), w)
	})

	app.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		pages.LogInPage().Render(r.Context(), w)

	})

	conf := &oauth2.Config{
		RedirectURL:  cred.RedirectUri,
		ClientID:     cred.ClientId,
		ClientSecret: cred.ClientSecret,
		Scopes:       []string{auth.ScopeIdentify},
		Endpoint:     auth.Endpoint,
	}

	app.Get("/discord-oauth", sentryHandler.HandleFunc(func(w http.ResponseWriter, r *http.Request) {

		state, err := auth.StateGenerator()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			panic(fmt.Sprintf("Error generating state: %s", err))
		}

		http.Redirect(w, r, conf.AuthCodeURL(state), http.StatusTemporaryRedirect)

	}),
	)

	app.Get("/discord-redirect", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("After log in redirected from discord")

	})

	// app.Group("/user", func(r chi.Router) {

	// 	r.Use(handlers.ProtectedRoutes(r))

	// 	r.Get("/edit", func(w http.ResponseWriter, r *http.Request) {

	// 		pages.EditPage().Render(r.Context(), w)
	// 	})

	// 	r.Get("/account", func(w http.ResponseWriter, r *http.Request) {

	// 	})

	// 	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {

	// 	})

	// })

	log.Fatal(http.ListenAndServe(port, app))

	fmt.Println("Listening on 3000")
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
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
