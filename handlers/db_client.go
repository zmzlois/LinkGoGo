package handlers

import (
	"context"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/joho/godotenv"
	"github.com/zmzlois/LinkGoGo/ent"
)

type DatabaseCred struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
	SslMode  string
}

func DBCredentials() *DatabaseCred {

	err := godotenv.Load()

	if err != nil {
		panic("[Database] Error loading environment variables")
	}

	var cred *DatabaseCred = &DatabaseCred{
		Host:     "localhost",
		Port:     "5432",
		User:     os.Getenv("DB_USER"),
		Dbname:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SslMode:  "disable",
	}
	return cred
}

func DatabaseClient() {

	client, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=postgres dbname=world password=postgres123 sslmode=disable")

	if err != nil {
		log.Printf("failed opening connection to postgres: %v", err)
	}
	ctx := context.Background()
	defer client.Close()
	if err := client.Schema.Create(ctx); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}

}

// Example usage

// func main() {
// 	dbHandler := &handlers.DatabaseHandler{}

// 	ctx := context.Background()
// 	client, err := dbHandler.NewClient(ctx)
// 	if err != nil {
// 		log.Fatalf("failed to create database client: %v", err)
// 	}
// 	defer client.Close() // Close the client when you're done using it

// Now you can use the client to perform your queries
// For example:
// _, err := client.User.Create().SetUsername("example").Save(ctx)
// if err != nil {
//     log.Fatalf("failed to create user: %v", err)
// }
// }
