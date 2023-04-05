//go:generate go run github.com/Khan/genqlient

package client

import (
	"net/http"
	"net/url"

	"github.com/Khan/genqlient/graphql"
)

func NewClient(serverAddr *url.URL) graphql.Client {
	serverAddr.Path = "/graphql"
	return graphql.NewClient(serverAddr.String(), http.DefaultClient)
}
