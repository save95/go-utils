package ziputil

import (
	"archive/zip"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_testZipFile = "222357.zip"
)

func TestDecompressionBy(t *testing.T) {
	err := DecompressionBy(_testZipFile, func(file *zip.File) error {
		t.Log(file.Name)
		return nil
	})
	assert.Nil(t, err)
}
