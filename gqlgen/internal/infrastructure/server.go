package infrastructure

import (
	"net/http"

	"github.com/exepirit/go-graphql/gqlgen/internal/config"
	"github.com/exepirit/go-graphql/gqlgen/pkg/routing"
	"github.com/gin-gonic/gin"
)

func NewServer(cfg config.Config) (Server, error) {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	ginHandler := gin.Default()

	server := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: ginHandler,
	}

	return Server{
		Server: server,
	}, nil
}

type Server struct {
	*http.Server
}

func (srv Server) Bind(bindable routing.Bindable) {
	if bindable == nil {
		return
	}
	bindable.Bind(srv.Handler.(*gin.Engine))
}
