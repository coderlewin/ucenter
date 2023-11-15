package web

import (
	"github.com/cloudwego/hertz/pkg/route"
)

type Router interface {
	ConfigRoutes(h *route.RouterGroup)
}
