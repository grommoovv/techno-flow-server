package utils

import (
	"math/rand"
	"time"
)

func GenerageRandomDate(random *rand.Rand) time.Time {
	minValue := time.Now().Add(2 * time.Hour).Unix()
	maxValue := time.Now().Add(7 * 24 * time.Hour).Unix()
	delta := maxValue - minValue

	sec := random.Int63n(delta) + minValue
	return time.Unix(sec, 0)
}
