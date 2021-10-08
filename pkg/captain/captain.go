package captain

import "github.com/gin-gonic/gin"


type Captain struct {
	*gin.Engine
	router *gin.RouterGroup
}

type ICaptain interface {
	Build(caption *Captain)
}

func NewCaptain() *Captain {
	return &Captain{Engine: gin.New()}
}

func (c *Captain) Launch() {
	_ = c.Run(":8089")
}

func (c *Captain) Attach(f gin.HandlerFunc) *Captain{
	c.Use(f)
	return c
}

func (c *Captain) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) *Captain {
	c.router.Handle(httpMethod, relativePath, handlers...)
	return c
}

func (c *Captain) Mount(group string, ic ...ICaptain) *Captain{
	c.router = c.Group(group)
	for _,i :=range ic{
		i.Build(c)
	}
	return c
}