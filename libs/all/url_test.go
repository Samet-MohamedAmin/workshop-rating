package libs

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var testCasesGetIP = []struct {
	host string
	ip   string
}{
	{host: "bla:8080", ip: "bla"},
	{host: "192.168.1.43:8080", ip: "192.168.1.43"},
}

func TestGetIP(t *testing.T) {
	for _, tc := range testCasesGetIP {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{Host: tc.host}
		actual := GetIP(c)
		expected := tc.ip
		if actual != expected {
			t.Fatalf("getIP(), given %s,\nactual %s, expected %s", tc.host, actual, expected)
		}
	}
}
