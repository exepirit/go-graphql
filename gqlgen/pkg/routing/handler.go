package routing

import "github.com/gin-gonic/gin"

type bindableHandlerFunc struct {
	handler gin.HandlerFunc
}

func (b bindableHandlerFunc) Bind(router gin.IRouter) {
	router.Use(b.handler)
}

func HandlerFunc(handler gin.HandlerFunc) Bindable {
	return &bindableHandlerFunc{
		handler: handler,
	}
}
