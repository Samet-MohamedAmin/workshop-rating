package core

import (
	"net/http"

	. "rating/libs/all"
	service "rating/services/mongo"

	"github.com/gin-gonic/gin"
)

func IsSeen(c *gin.Context) {
	ip := GetIP(c)

	msg := "NOT SEEN"
	if _, ok := service.GetOne(ip); ok {
		msg = "SEEN"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
