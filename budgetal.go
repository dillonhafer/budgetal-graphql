package main

import "os"

type Budgetal struct {
	env string
}

func (b *Budgetal) extractEnv() {
	if os.Getenv("BUDGETAL_ENV") == "" {
		b.env = "development"
	} else {
		b.env = os.Getenv("BUDGETAL_ENV")
	}
}

func (b *Budgetal) production() bool {
	return b.env == "production"
}
