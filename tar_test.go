package zdpgo_tar

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"path"
	"testing"
)

/*
@Time : 2022/6/2 16:02
@Author : 张大鹏
@File : tar_test.go
@Software: Goland2021.3.1
@Description:
*/

// 测试压缩文件
func TestTar_TarGz(t *testing.T) {
	tar := New(zdpgo_log.NewWithDebug(true, "log.log"))
	dirPath := "C:\\projects\\go\\xjaq\\sev_versions\\1.0.0.5\\"
	dirList := []string{
		"xmsev_1000005_windows_amd64",
		"xmsev_1000005_linux_aarch64",
		"xmsev_1000005_mac_darwin64",
		"xmsev_1000005_linux_amd64",
	}
	for _, dir := range dirList {
		dest := path.Join(dirPath, dir)
		err := tar.TarGz(dest, dest+".tar.gz")
		if err != nil {
			panic(err)
		}
	}

	dirPath = "C:\\projects\\go\\xjaq\\sev_assistant_versions\\1.0.0.5\\"
	dirList = []string{
		"xmsev_assistant_1000005_windows_amd64",
		"xmsev_assistant_1000005_linux_aarch64",
		"xmsev_assistant_1000005_mac_darwin64",
		"xmsev_assistant_1000005_linux_amd64",
	}
	for _, dir := range dirList {
		dest := path.Join(dirPath, dir)
		err := tar.TarGz(dest, dest+".tar.gz")
		if err != nil {
			panic(err)
		}
	}

}

// 测试压缩指定文件夹下的指定子文件夹
func TestTar_TarGzDirFiles(t *testing.T) {
	tar := New(zdpgo_log.NewWithDebug(true, "log.log"))
	dirPath := "C:\\projects\\go\\xjaq\\sev_versions\\1.0.0.5\\"
	dirList := []string{
		"xmsev_1000005_windows_amd64",
		"xmsev_1000005_linux_aarch64",
		"xmsev_1000005_mac_darwin64",
		"xmsev_1000005_linux_amd64",
	}
	err := tar.TarGzDirFiles(dirPath, dirList)
	if err != nil {
		panic(err)
	}
}

// 测试压缩指定文件夹下的所有子文件夹
func TestTar_TarGzDirAllFiles(t *testing.T) {
	tar := New(zdpgo_log.NewWithDebug(true, "log.log"))
	dirPath := "C:\\projects\\go\\xjaq\\sev_versions\\1.0.0.5\\"
	err := tar.TarGzDirAllFiles(dirPath)
	if err != nil {
		panic(err)
	}
}

// 测试解压
func TestTar_UnTarGz(t *testing.T) {
	tar := New(zdpgo_log.NewWithDebug(true, "log.log"))
	dirPath := "C:\\projects\\go\\xjaq\\sev_versions\\1.0.0.5\\test\\xmsev_1000005_linux_aarch64.tar.gz"
	err := tar.UnTarGz(dirPath, "C:\\projects\\go\\xjaq\\sev_versions\\1.0.0.5\\test\\xmsev_1000005_linux_aarch64")
	if err != nil {
		panic(err)
	}
}

// 测试解压到同级目录
func TestTar_UnTarGzToSameDir(t *testing.T) {
	tar := New(zdpgo_log.NewWithDebug(true, "log.log"))
	dirPath := "C:\\projects\\go\\xjaq\\sev_versions\\1.0.0.5\\test\\xmsev_1000005_linux_aarch64.tar.gz"
	err := tar.UnTarGzToSameDir(dirPath)
	if err != nil {
		panic(err)
	}
}

// 测试解压到同级目录并删除压缩文件
func TestTar_UnTarGzToSameDirAndDelete(t *testing.T) {
	tar := New(zdpgo_log.NewWithDebug(true, "log.log"))
	dirPath := "C:\\projects\\go\\xjaq\\sev_versions\\1.0.0.5\\test\\xmsev_1000005_linux_aarch64.tar.gz"
	err := tar.UnTarGzToSameDirAndDelete(dirPath)
	if err != nil {
		panic(err)
	}
}
