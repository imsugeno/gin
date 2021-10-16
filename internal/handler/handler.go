package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewGinEngin() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	return r
}
