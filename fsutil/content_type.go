package fsutil

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/save95/xerror"
)

// MIMEType 获得文件的 Content-Type，是使用 ParseFileContentType
// Deprecated
func MIMEType(filename string) (string, error) {
	return ParseFileContentType(filename)
}

func ParseFileContentType(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", xerror.Wrap(err, "file open failed")
	}
	defer func() {
		_ = f.Close()
	}()

	return ParseContentType(f)
}

func ParseContentType(reader io.Reader) (string, error) {
	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)
	if _, err := reader.Read(buffer); err != nil {
		return "", xerror.Wrap(err, "file read failed")
	}

	return http.DetectContentType(buffer), nil
}

func Extension(filename string) (string, error) {
	ct, err := ParseFileContentType(filename)
	if nil != err {
		return "", err
	}

	return ContentType2Ext(ct)
}

func ContentType2Ext(contentType string) (string, error) {
	// 如果 text/plain; charset=utf-16be
	// 只返回 text/plain
	key := strings.ToLower(strings.Split(contentType, ";")[0])
	if ext, ok := mimeTypeMapping[key]; ok {
		return ext, nil
	} else {
		return "", xerror.Errorf("no support content-type(%s)", contentType)
	}
}
