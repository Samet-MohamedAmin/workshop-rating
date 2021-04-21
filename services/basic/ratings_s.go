package rating

import (
	models "rating/core/models"
)

var rs = models.Ratings{}

const (
	errEmptyParam = "%s param is empty"
	errBadParam   = "%s param must be an integer value between 0 and 5 inclusive"
)

func EmptyRatings() {
	rs = models.Ratings{}
}

// TODO: should return []models.RatingItem
func GetRatings() models.Ratings {
	return rs
}

func AddRating(ip string, rating int) {
	rs[ip] = rating
}

func GetOne(ip string) (ri models.RatingItem, ok bool) {
	r, ok := rs[ip]
	ri = models.RatingItem{}
	if ok {
		ri.Ip = ip
		ri.Rating = r
	}
	return
}

func GetAverage() float32 {
	rs = GetRatings()
	return rs.GetAverage()
}
