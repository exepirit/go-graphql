//go:generate go run github.com/99designs/gqlgen generate

package main

import (
	"github.com/exepirit/go-graphql/gqlgen/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.Module).Run()
}
