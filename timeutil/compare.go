package timeutil

import (
	"time"
)

func Greatest(x time.Time, times ...time.Time) time.Time {
	greatest := x
	for _, t := range times {
		if greatest.Before(t) {
			greatest = t
		}
	}

	return greatest
}

func Lowest(x time.Time, times ...time.Time) time.Time {
	lowest := x
	for _, t := range times {
		if lowest.After(t) {
			lowest = t
		}
	}

	return lowest
}
