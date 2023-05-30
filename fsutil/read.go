package fsutil

import (
	"encoding/base64"
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func ReadFileToBase64(path string) (string, error) {
	data, err := ReadFile(path)
	if nil != err {
		return "", err
	}

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	return sEnc, nil
}
