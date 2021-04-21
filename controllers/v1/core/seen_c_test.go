package core

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	service "rating/services/mongo"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testCasesIsSeen = []struct {
	host string
	ip   string
}{
	{host: "bla:8080", ip: "bla"},
	{host: "192.168.1.43:8080", ip: "192.168.1.43"},
}

func TestIsSeen(t *testing.T) {
	// NOT SEEN
	for _, tc := range testCasesIsSeen {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{Host: tc.host}

		service.EmptyCollection()
		IsSeen(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var got gin.H
		err := json.Unmarshal(w.Body.Bytes(), &got)

		if err != nil {
			t.Fatal(err)
		}

		want := gin.H{"message": "NOT SEEN"}
		assert.Equal(t, want, got)
	}

	// SEEN
	for _, tc := range testCasesIsSeen {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = &http.Request{Host: tc.host}

		service.EmptyCollection()
		service.AddOne(tc.ip, 0)
		IsSeen(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var got gin.H
		err := json.Unmarshal(w.Body.Bytes(), &got)

		if err != nil {
			t.Fatal(err)
		}

		want := gin.H{"message": "SEEN"}
		assert.Equal(t, want, got)
	}
}
