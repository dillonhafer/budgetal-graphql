package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/handler"
	"github.com/jackc/pgx"
)

var conn *pgx.Conn

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	bindAddress := os.Getenv("BIND_ADDRESS")
	if bindAddress == "" {
		bindAddress = "127.0.0.1"
	}
	serveAddress := bindAddress + ":" + port

	var err error
	conn, err = pgx.Connect(extractConfig())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	schema := NewSchema()
	handle := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// serve HTTP
	printStartup(serveAddress)
	http.Handle("/", handle)
	http.ListenAndServe(serveAddress, nil)
}

func printStartup(serveAddress string) {
	println("=> Booting Budgetal")
	println("=> Application starting in development on http://" + serveAddress)
	println("=> Ctrl-C to shutdown server")
}

func extractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		database_url = "postgres://" + os.Getenv("USER") + "@localhost:5432/budgetal_development"
	}

	config, err := pgx.ParseConnectionString(database_url)
	if err != nil {
		log.Fatalf("failed to parse DATABASE_URL, error: %v", err)
	}

	return config
}
