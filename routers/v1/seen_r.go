package v1

import (
	controller "rating/controllers/v1/core"

	"github.com/gin-gonic/gin"
)

func SetSeenRoutes(router *gin.Engine) {
	router.GET("/seen", controller.IsSeen)
}
