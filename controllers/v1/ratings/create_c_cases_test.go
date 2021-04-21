package ratings

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var testCasesIsSeen = []struct {
	host string
	ip   string
}{
	{host: "bla:8080", ip: "bla"},
	{host: "192.168.1.43:8080", ip: "192.168.1.43"},
}

var testCasesAddRatingNotValid = []struct {
	urlValues   url.Values
	status      int
	response    gin.H
	description string
}{
	{
		urlValues:   url.Values{},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errEmptyParam, "rating")},
		description: "no params",
	},
	{
		urlValues:   url.Values{"rating": {}},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errEmptyParam, "rating")},
		description: "not decalerd param",
	},
	{
		urlValues:   url.Values{"rating": {""}},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errEmptyParam, "rating")},
		description: "empty param",
	},
	{
		urlValues:   url.Values{"rating": {"a"}},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errBadParam, "rating")},
		description: "not digit",
	},
	{
		urlValues:   url.Values{"rating": {"z"}},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errBadParam, "rating")},
		description: "not digit",
	},
	{
		urlValues:   url.Values{"rating": {"7a"}},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errBadParam, "rating")},
		description: "not digit",
	},
	{
		urlValues:   url.Values{"rating": {"6"}},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errBadParam, "rating")},
		description: "greater than 5",
	},
	{
		urlValues:   url.Values{"rating": {"-1"}},
		status:      http.StatusBadRequest,
		response:    gin.H{"message": fmt.Sprintf(errBadParam, "rating")},
		description: "lower than 0",
	},
}
