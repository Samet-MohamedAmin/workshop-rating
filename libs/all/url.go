package libs

import (
	"github.com/gin-gonic/gin"
)

// GetIP gets a requests IP address
func GetIP(c *gin.Context) string {
	// TODO: better approach https://www.golanglearn.com/get-ip-address-using-host-header-using-golang/
	var ip string
	forwarded := c.Request.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ip = forwarded
	} else {
		ip = c.Request.RemoteAddr
	}
	return ip
}
