package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnake(t *testing.T) {
	strs := map[string]string{
		"view_rank": "view_rank",
		"viewRank":  "view_rank",
	}

	for str, target := range strs {
		assert.Equal(t, Snake(str), target)
	}
}
