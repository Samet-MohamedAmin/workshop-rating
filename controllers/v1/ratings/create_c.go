package ratings

import (
	"fmt"
	"net/http"
	. "rating/libs/all"
	service "rating/services/mongo"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	errEmptyParam = "%s param is empty"
	errBadParam   = "%s param must be an integer value between 0 and 5 inclusive"
)

type PostBody struct {
	Rating int `json:"rating,omitempty"`
}

func AddRating(c *gin.Context) {
	ip := GetIP(c)

	var body PostBody
	if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf(errBadParam, "rating"),
		})
		return
	}

	service.AddOne(ip, body.Rating)

	c.JSON(http.StatusOK, gin.H{
		"ip":     ip,
		"rating": body.Rating,
	})
}
