package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetPingRoutes(router *gin.Engine) {
	// start the server
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
