package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/graphql-go/handler"
	"github.com/jackc/pgx"
)

var conn *pgx.Conn

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

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
	http.Handle("/", handle)
	http.ListenAndServe("127.0.0.1:"+port, nil)
}

func extractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		database_url = "postgres://" + os.Getenv("USER") + "@localhost:5432/budgetal_development"
	}

	url, err := url.Parse(database_url)
	if err != nil {
		log.Fatalf("failed to parse DATABASE_URL, error: %v", err)
	}

	if url.Scheme != "postgres" {
		log.Fatalf("DATABASE_URL scheme must be postgres://")
	}

	database := strings.Replace(url.RequestURI(), "/", "", -1)
	host, port_string, _ := net.SplitHostPort(url.Host)
	port, err := strconv.ParseInt(port_string, 10, 16)
	if err != nil {
		log.Fatalf("failed to parse database port, error: %v", err)
	}

	user := url.User.Username()
	password, _ := url.User.Password()

	config.Host = host
	config.Port = uint16(port)
	config.Database = database
	config.User = user
	config.Password = password

	return config
}
