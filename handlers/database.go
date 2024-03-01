package handlers

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/zmzlois/LinkGoGo/ent"
)

func DatabaseClient(ctx context.Context) *ent.Client {

	client, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=postgres dbname=linkgogo password=postgres sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}

	return client
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
