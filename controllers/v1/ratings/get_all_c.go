package ratings

import (
	"net/http"
	service "rating/services/mongo"

	"github.com/gin-gonic/gin"
)

func GetRatings(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAll())
}
