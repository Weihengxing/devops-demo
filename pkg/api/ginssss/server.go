package ginssss

import (
	"github.com/Weihengxing/devops-demo/pkg/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

var _ api.IServer = &Gin{}

type Gin struct {
	*gin.Engine
}



func NewGin() *Gin{
	route:=gin.Default()
	return &Gin{route}
}

func (g *Gin) Listen(addr string) {
	panic("implement me")
}

func (g *Gin) Run() error {
	panic("implement me")
}

func (g *Gin) Register(s string, handler http.Handler) error {
	panic("implement me")
}

