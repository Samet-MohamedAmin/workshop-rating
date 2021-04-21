package v1

import (
	ratings "rating/controllers/v1/ratings"

	"github.com/gin-gonic/gin"
)

func SetRatingRoutes(router *gin.Engine) {
	router.POST("/ratings", ratings.AddRating)
	router.GET("/ratings", ratings.GetRatings)
	router.GET("/ratings/average", ratings.GetAverage)
}
