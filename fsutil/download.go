package fsutil

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Download 下载远程文件
func Download(filename, url string) error {
	// todo 检查url

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("remote response failed: status code = %d", resp.StatusCode)
	}

	// 创建文件
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	// 保存下载文件
	_, err = io.Copy(out, resp.Body)
	return err
}
