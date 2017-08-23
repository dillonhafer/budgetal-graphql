package main

import (
	"net/http"
	"os"

	"github.com/graphql-go/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
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
