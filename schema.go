package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func NewSchema() graphql.Schema {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "kevin", nil
			},
		},
		"allUsers": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				err := listUsers()
				return "kevin", err
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}

func listUsers() error {
	rows, _ := conn.Query("select email, first_name from users")

	for rows.Next() {
		var email string
		var first_name string
		err := rows.Scan(&email, &first_name)
		if err != nil {
			return err
		}
		fmt.Printf("%s. %s\n", email, first_name)
	}

	return rows.Err()
}
