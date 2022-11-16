package fsutil

import (
	"mime"
	"testing"
)

func TestExtension(t *testing.T) {
	mts := map[string]string{
		"application/octet-stream":     "",
		"text/plain; charset=utf-16be": "",
		"image/bmp":                    "",
		"image/jpeg":                   "",
	}

	for mt := range mts {
		es, err := mime.ExtensionsByType(mt)
		t.Log(mt, " -> ", es, err)
	}
}
