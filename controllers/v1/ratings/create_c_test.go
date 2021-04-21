package ratings

import (
	"testing"
)

func TestAddRatingNotValid(t *testing.T) {
	// host := "192.168.1.43:8080"

	// for _, tc := range testCasesAddRatingNotValid {
	// 	w := httptest.NewRecorder()
	// 	c, _ := gin.CreateTestContext(w)
	// 	c.Request = &http.Request{Host: host}
	// 	c.Request.PostForm = tc.urlValues

	// 	service.EmptyRatings()
	// 	AddRating(c)

	// 	assert.Equal(t, tc.status, w.Code)

	// 	var got gin.H
	// 	err := json.Unmarshal(w.Body.Bytes(), &got)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	want := tc.response
	// 	assert.Equal(t, want, got)
	// }
}

func TestAddRatingValid(t *testing.T) {
	// ip := "192.168.1.43"
	// host := ip + ":8080"

	// for i := 0; i <= 5; i++ {
	// 	w := httptest.NewRecorder()
	// 	c, _ := gin.CreateTestContext(w)
	// 	c.Request = &http.Request{Host: host}
	// 	c.Request.PostForm = url.Values{"rating": {fmt.Sprint(i)}}

	// 	service.EmptyRatings()
	// 	AddRating(c)

	// 	assert.Equal(t, http.StatusOK, w.Code)

	// 	var got gin.H
	// 	err := json.Unmarshal(w.Body.Bytes(), &got)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	rs := service.GetRatings()
	// 	want := gin.H{
	// 		"ip":     ip,
	// 		"rating": float64(rs[ip]),
	// 	}
	// 	assert.Equal(t, want, got)
	// }
}
