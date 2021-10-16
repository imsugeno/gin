package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func bindRouteOfHello(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
}
