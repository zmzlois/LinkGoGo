package main

// TODO: middleware if you don't have any cookies "github or discord", bounce back to login page
// TODO: how are you getting the cookie

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	m "github.com/zmzlois/LinkGoGo/pkg/middleware"
	"github.com/zmzlois/LinkGoGo/pkg/monitor"
	"github.com/zmzlois/LinkGoGo/pkg/router"
	"github.com/zmzlois/LinkGoGo/pkg/utils"
)

func main() {

	port := utils.Config("PORT")

	fmt.Printf("Port is trying to listening on %s\n\n", port)

	// Initialize sentry
	monitor.SentryInit()
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	app := chi.NewRouter()

	app.Use(middleware.Logger)

	app.Use(monitor.SentryHandler.Handle)
	app.Use(m.PublicMiddleware)

	// Serving static files: Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "/web/assets"))
	utils.FileServer(app, "/", filesDir)

	// set up all the routers
	router.SetupRouter(app)

	log.Fatal(http.ListenAndServe(port, app))

	fmt.Println("Listening on 3000")
}
