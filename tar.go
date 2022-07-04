package zdpgo_tar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
@Time : 2022/6/2 15:07
@Author : 张大鹏
@File : tar.go
@Software: Goland2021.3.1
@Description:
*/

type Tar struct {
	Config *Config
}

func New() *Tar {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *Tar {
	t := &Tar{}

	// 配置
	t.Config = config

	// 返回
	return t
}

// TarGz 压缩.tar.gz格式
func (t *Tar) TarGz(srcDirPath string, destFilePath string) error {
	fw, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer fw.Close()

	// 创建gzip写入器
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 创建tar写入器
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// 检查是文件夹还是文件
	f, err := os.Open(srcDirPath)
	if err != nil {
		return err
	}

	// 获取文件信息
	fi, err := f.Stat()
	if err != nil {
		return err
	}

	// 对文件夹和文件做不同的处理
	if fi.IsDir() {
		err = t.TarGzDir(srcDirPath, path.Base(srcDirPath), tw)
		if err != nil {
			return err
		}
	} else {
		err = t.TarGzFile(srcDirPath, fi.Name(), tw, fi)
		if err != nil {
			return err
		}
	}

	// 返回
	return nil
}

// TarGzDirFiles 压缩指定文件夹下的指定子文件夹
func (t *Tar) TarGzDirFiles(dirPath string, files []string) error {
	// 压缩指定文件夹下的每个子文件夹
	for _, dir := range files {
		dest := path.Join(dirPath, dir)
		err := t.TarGz(dest, dest+".tar.gz")
		if err != nil {
			return err
		}
	}

	// 返回
	return nil
}

// TarGzDirAllFiles 压缩指定文件夹下的所有子文件夹
func (t *Tar) TarGzDirAllFiles(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// 压缩所有的文件夹
	for _, file := range files {
		if file.IsDir() {
			dest := path.Join(dirPath, file.Name())
			err = t.TarGz(dest, dest+".tar.gz")
			if err != nil {
				return err
			}
		}
	}

	// 返回
	return nil
}

// TarGzDir 压缩文件夹
func (t *Tar) TarGzDir(srcDirPath string, recPath string, tw *tar.Writer) error {
	// 打开文件夹
	dir, err := os.Open(srcDirPath)
	if err != nil {
		return err
	}
	defer dir.Close()

	// 读取文件
	fis, err := dir.Readdir(0)
	if err != nil {
		return err
	}

	// 处理每个文件
	for _, fi := range fis {
		// 追加路径
		curPath := srcDirPath + "/" + fi.Name()
		// 如果是文件夹，递归压缩
		if fi.IsDir() {
			err = t.TarGzDir(curPath, recPath+"/"+fi.Name(), tw)
			if err != nil {
				return err
			}
		}
		// 压缩文件
		err = t.TarGzFile(curPath, recPath+"/"+fi.Name(), tw, fi)
		if err != nil {
			return err
		}
	}

	// 返回
	return nil
}

// TarGzFile 压缩文件
func (t *Tar) TarGzFile(srcFile string, recPath string, tw *tar.Writer, fi os.FileInfo) error {
	// 如果是文件夹
	if fi.IsDir() {
		hdr := new(tar.Header)
		hdr.Name = recPath + "/"
		hdr.Typeflag = tar.TypeDir
		hdr.Size = 0
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()
		err := tw.WriteHeader(hdr)
		if err != nil {
			return err
		}
	} else {
		// 打开文件
		fr, err := os.Open(srcFile)
		if err != nil {
			return err
		}
		defer fr.Close()

		// 创建头部
		hdr := new(tar.Header)
		hdr.Name = recPath
		hdr.Size = fi.Size()
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		// 写入头部
		err = tw.WriteHeader(hdr)
		if err != nil {
			return err
		}

		// 写入数据
		_, err = io.Copy(tw, fr)
		if err != nil {
			return err
		}
	}

	// 返回
	return nil
}

// UnTarGz 解压缩.tar.gz文件
func (t *Tar) UnTarGz(srcFilePath string, destDirPath string) error {
	if destDirPath == "" {
		return errors.New("指定目标目录不能为空")
	}

	err := os.Mkdir(destDirPath, os.ModePerm)
	if err != nil {
		return err
	}

	// 打开源文件
	fr, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer fr.Close()

	// 创建gzip读取器
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}

	// 创建tar读取器
	tr := tar.NewReader(gr)

	for {
		var hdr *tar.Header
		hdr, err = tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if hdr.Typeflag != tar.TypeDir {
			err = os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			if err != nil {
				return err
			}

			// 写入文件数据
			var fw *os.File
			fw, err = os.Create(destDirPath + "/" + hdr.Name)
			if err != nil {
				return err
			}

			// 复制文件
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}

	// 返回
	return nil
}

func (t *Tar) DeleteDirAll(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".tar.gz") {
			err = os.Remove(filepath.Join(dirPath, f.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
func (t *Tar) UnTarGzDir(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".tar.gz") {
			dirName := strings.Replace(f.Name(), ".tar.gz", "", 1)

			// auto rename
			targetName := filepath.Join(dirPath, dirName)
			if t.IsExists(targetName) {
				var (
					tmpName = targetName
					count   = 1
				)
				for {
					targetName += fmt.Sprintf("(%d)", count)
					if !t.IsExists(tmpName) {
						break
					}
					tmpName += fmt.Sprintf("(%d)", count)
					count++
				}
			}

			// execute un compress
			err = t.UnTarGz(filepath.Join(dirPath, f.Name()), targetName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Tar) IsExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// UnTarGzToSameDir 解压.tar.gz压缩包到该文件所在目录
func (t *Tar) UnTarGzToSameDir(srcFilePath string) error {
	dirPath, _ := filepath.Split(srcFilePath)

	// 打开源文件
	fr, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer fr.Close()

	// 创建gzip读取器
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}

	// 创建tar读取器
	tr := tar.NewReader(gr)

	for {
		var hdr *tar.Header
		hdr, err = tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if hdr.Typeflag != tar.TypeDir {
			err = os.MkdirAll(dirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			if err != nil {
				return err
			}

			// 写入文件数据
			var fw *os.File
			fw, err = os.Create(dirPath + "/" + hdr.Name)
			if err != nil {
				return err
			}

			// 复制文件
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}

	// 返回
	return nil
}

// UnTarGzToSameDirAndDelete 解压到同级目录，并删除原来的压缩包
func (t *Tar) UnTarGzToSameDirAndDelete(srcFilePath string) error {
	// 解压
	err := t.UnTarGzToSameDir(srcFilePath)
	if err != nil {
		return err
	}

	// 删除
	err = os.RemoveAll(srcFilePath)
	if err != nil {
		return err
	}

	// 返回
	return nil
}
