package models

type Ratings map[string]int

func (rs Ratings) GetAverage() float32 {
	if rs.isEmpty() {
		return 0
	}

	var sum, users float32 = 0, 0
	for _, r := range rs {
		sum += float32(r)
		users++
	}
	return sum / users
}

func (rs Ratings) isEmpty() bool {
	return len(rs) == 0
}
