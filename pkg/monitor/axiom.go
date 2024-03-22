package monitor

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/axiomhq/axiom-go/axiom"
	"github.com/axiomhq/axiom-go/axiom/ingest"
	"github.com/zmzlois/LinkGoGo/pkg/utils"
)

type Logger struct {
	client *axiom.Client
}

type LogInput struct {
	Message    string `json:"message"`
	Attributes string `json:"attributes"`
	Resources  interface{}
}

func filePath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting file path for OTel")
	}

	folderName := filepath.Base(wd)
	return folderName
}

func NewLogger() *Logger {
	client, err := axiom.NewClient(
		axiom.SetPersonalTokenConfig(utils.Config("AXIOM_TOKEN"), utils.Config("AXIOM_ORG_ID")),
	)
	if err != nil {
		log.Fatal(err)
	}
	return &Logger{client: client}
}

func (l *Logger) Info(ctx context.Context, input *LogInput) {
	event := []axiom.Event{
		{
			"level":               "info",
			ingest.TimestampField: time.Now(),
			"caller":              filePath(),
			"attributes":          input.Attributes,
			"message":             input.Message,
			"resources":           input.Resources,
		},
	}

	_, err := l.client.IngestEvents(ctx, "linkgogo", event)
	if err != nil {
		log.Fatalln(err)
	}
}

func (l *Logger) Debug() {}

func (l *Logger) Warn() {}

func (l *Logger) Error() {}

// suppose to behave the same as panic in development and error in production
func (l *Logger) DPanic() {
}

func (l *Logger) Panic() {}

func (l *Logger) Fatal() {}

// authorid, link-slug, time
func (l *Logger) LinkAnalytic() {}
