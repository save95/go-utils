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
		"audio/mp3":                    "",
	}

	for mt := range mts {
		es, err := mime.ExtensionsByType(mt)
		ext, err2 := ContentType2Ext(mt)
		t.Log(mt, " -> ", es, err, " ==> ", ext, err2)
	}
}
