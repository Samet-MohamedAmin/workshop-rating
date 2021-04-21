package v1

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	SetPingRoutes(router)
	SetSeenRoutes(router)
	SetRatingRoutes(router)
}
