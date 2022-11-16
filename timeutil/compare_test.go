package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGreatest(t *testing.T) {
	a := time.Date(2022, 1, 2, 12, 1, 32, 21, time.Local)
	times := []time.Time{
		time.Date(2022, 1, 2, 11, 0, 0, 0, time.Local),
		time.Date(2022, 1, 2, 12, 1, 32, 20, time.Local),
		time.Date(2022, 1, 2, 12, 1, 30, 21, time.Local),
	}

	tps := make([]time.Time, 0)
	for _, t2 := range times {
		tps = append(tps, t2)
	}

	assert.Equal(t, a, Greatest(a, tps...))
}
