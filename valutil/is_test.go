package valutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNil(t *testing.T) {
	var (
		//a bool
		//b uint
		//c int
		//d struct{}
		//e string
		e interface{}
		f *string
		g *bool
		h *struct{}
	)

	e = f
	nils := []interface{}{
		e, f, g, h,
	}
	for _, val := range nils {
		t.Logf("%#v", val)
		assert.True(t, IsNil(val))
	}

	boolVal := true
	intVal := 1000
	e = boolVal
	noNils := []interface{}{
		true, false, 0, uint(0), struct{}{},
		&boolVal, &intVal, e,
	}

	for _, val := range noNils {
		assert.False(t, IsNil(val))
	}
}
