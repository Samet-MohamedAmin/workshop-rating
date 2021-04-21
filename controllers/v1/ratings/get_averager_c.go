package ratings

import (
	"net/http"
	service "rating/services/mongo"

	"github.com/gin-gonic/gin"
)

func GetAverage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"average": service.GetAverage(),
	})
}
