package monitor

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/joho/godotenv"
)

type handler struct{}

func (h *handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
		hub.WithScope(func(scope *sentry.Scope) {
			scope.SetExtra("unwantedQuery", "someQueryDataMaybe")
			hub.CaptureMessage("User provided unwanted query string, but we recovered just fine")
		})
	}
	rw.WriteHeader(http.StatusOK)
}

func enhanceSentryEvent(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		handler(rw, r)
	}
}

func SentryInit() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sentryDSN := os.Getenv("SENTRY_DSN")

	if len(sentryDSN) == 0 {
		panic("SENTRY_DSN is required for sentry to work")
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					fmt.Printf("Request URL: %s", req.URL)
				}
			}

			return event
		},
		Debug:            true,
		AttachStacktrace: true,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		panic(fmt.Sprintf("[Sentry error]: %s", err))
	}

}

var Sh = sentryhttp.New(sentryhttp.Options{
	Repanic:         true,
	WaitForDelivery: true,
	Timeout:         2 * time.Second,
})
