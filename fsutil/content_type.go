package fsutil

import (
	"net/http"
	"os"

	"github.com/save95/xerror"
)

func MIMEType(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", xerror.Wrap(err, "file open failed")
	}
	defer func() {
		_ = f.Close()
	}()

	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)
	if _, err := f.Read(buffer); err != nil {
		return "", xerror.Wrap(err, "file read failed")
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

//func Extension(filename string) ([]string, error) {
//	ct, err := MIMEType(filename)
//	if nil != err {
//		return nil, err
//	}
//
//	return mime.ExtensionsByType(ct)
//}
