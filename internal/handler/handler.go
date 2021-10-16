package handler

import (
	"github.com/gin-gonic/gin"
)

func NewGinEngin() *gin.Engine {
	r := gin.Default()

	public := r.Group("/public")
	{
		bindRouteOfHello(public.Group("/hello"))
	}
	return r
}
