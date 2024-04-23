package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/zmzlois/LinkGoGo/ent"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
	SslMode  string
}

// var db *hdl.Database = hdl.DBInit(&hdl.Database{
// 	Host:     hdl.DBCredentials().Host,
// 	Port:     hdl.DBCredentials().Port,
// 	User:     hdl.DBCredentials().User,
// 	Dbname:   hdl.DBCredentials().Dbname,
// 	Password: hdl.DBCredentials().Password,
// 	SslMode:  hdl.DBCredentials().SslMode,
// })

func DBCredentials() *Database {

	var cred *Database = &Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Dbname:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SslMode:  "disable",
	}
	return cred
}

func (db *Database) checkStructErrors() {
	// Make sure the user has provided
	// a valid client id
	if len(db.Host) < 1 {
		panic("Database connection Error: invalid Hostname in DBClient (ClientID: string)")
	}
	// Make sure the user has provided
	// a valid client secret
	if len(db.Port) < 1 {
		panic("Database connection Error: invalid Port in DBClient ")
	}
	// Make sure the user has provided
	// a valid redirect uri
	if len(db.Password) < 1 {
		panic("Database connection Error: invalid Password in DBClient ")
	}
	// Make sure the user has provided
	// a sufficient number of scopes
	if len(db.User) < 1 {
		panic("Database connection Error: invalid Username in DBClient")
	}
}

var Dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", DBCredentials().Host, DBCredentials().Port, DBCredentials().User, DBCredentials().Dbname, DBCredentials().Password)

func Open(dsn string) *ent.Client {

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	// Cache
	driver := entcache.NewDriver(
		drv,
		entcache.TTL(time.Second),
		entcache.Levels(entcache.NewLRU(128)),
	)
	return ent.NewClient(ent.Driver(driver))
}

func Connection() *ent.Client {
	client := Open(Dsn)
	// Run the migrations.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("failed creating schema resources: %v", err)
	}
	return client

}
