package routing

import "github.com/gin-gonic/gin"

func Route(prefix string, bind ...Bindable) Bindable {
	return &route{
		prefix: prefix,
		sub:    bind,
	}
}

type route struct {
	prefix string
	sub    []Bindable
}

func (r route) Bind(router gin.IRouter) {
	router = router.Group(r.prefix)
	for _, subRoute := range r.sub {
		subRoute.Bind(router)
	}
}

func Union(bind ...Bindable) Bindable {
	return &union{
		sub: bind,
	}
}

type union struct {
	sub []Bindable
}

func (u union) Bind(router gin.IRouter) {
	for _, subRoute := range u.sub {
		subRoute.Bind(router)
	}
}
