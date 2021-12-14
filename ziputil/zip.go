package ziputil

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CompressPath 压缩一个指定目录
// 将 src 目录打包成 dst 文件
func CompressPath(dst, src string) error {
	// 创建准备写入的文件
	fw, err := os.Create(dst)
	defer func() {
		_ = fw.Close()
	}()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zipW := zip.NewWriter(fw)
	defer func() {
		_ = zipW.Close()
	}()

	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		// 通过文件信息，创建 zip 的文件信息
		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}

		// 替换文件名中的相对路径
		fh.Name = strings.TrimPrefix(strings.TrimPrefix(path, src), string(filepath.Separator))

		// 这步开始没有加，会发现解压的时候说它不是个目录
		if fi.IsDir() {
			fh.Name += "/"
		}

		// 写入文件信息，并返回一个 Write 结构
		w, err := zipW.CreateHeader(fh)
		if err != nil {
			return
		}

		// 检测，如果不是标准文件就只写入头信息，不写入文件数据到 w
		// 如目录，也没有数据需要写
		if !fh.Mode().IsRegular() {
			return nil
		}

		// 打开要压缩的文件
		fr, err := os.Open(path)
		defer func() {
			_ = fr.Close()
		}()
		if err != nil {
			return err
		}

		// 将打开的文件 Copy 到 w
		if _, err := io.Copy(w, fr); err != nil {
			return err
		}

		// 输出压缩的内容
		//fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)

		return nil
	})
}

// DecompressionBy 通过自定义函数解压文件
// 将在 fun 函数返回 error 时中断解压，并返回 error
func DecompressionBy(srcFile string, fun func(file *zip.File) error) error {
	zr, err := zip.OpenReader(srcFile)
	defer func() {
		_ = zr.Close()
	}()
	if err != nil {
		return err
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		if err := fun(file); nil != err {
			return err
		}
	}
	return nil
}

// Decompression 解压 zip 包
// 将 srcFile 压缩包解压至 dstDir
func Decompression(srcFile, dstDir string) error {
	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dstDir != "" {
		if err := os.MkdirAll(dstDir, 0755); err != nil {
			return err
		}
	}

	return DecompressionBy(srcFile, func(file *zip.File) error {
		filename := filepath.Join(dstDir, file.Name)

		// 如果是目录，就创建目录
		// 因为是目录，跳过当前循环，因为后面都是文件的处理
		if file.FileInfo().IsDir() {
			return os.MkdirAll(filename, file.Mode())
		}

		return WriteFile(file, filename)
	})
}

// WriteFile 将 zip 中文件写入磁盘
func WriteFile(file *zip.File, filename string) error {
	// 获取到 Reader
	fr, err := file.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = fr.Close()
	}()

	// 创建要写出的文件对应的 Write
	fw, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer func() {
		_ = fw.Close()
	}()

	if _, err = io.Copy(fw, fr); err != nil {
		return err
	}

	return nil
}
