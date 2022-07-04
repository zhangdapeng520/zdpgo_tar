package main

import (
	"github.com/zhangdapeng520/zdpgo_tar"
)

/*
@Time : 2022/6/17 15:06
@Author : 张大鹏
@File : main.go.go
@Software: Goland2021.3.1
@Description:
*/

func main() {
	tar := zdpgo_tar.New()
	dirPath := "examples/delete_dir_all"

	// compress
	err := tar.TarGzDirAllFiles(dirPath)
	if err != nil {
		panic(err)
	}

	// delete
	err = tar.DeleteDirAll(dirPath)
	if err != nil {
		panic(err)
	}

	// uncompress
	err = tar.UnTarGzDir(dirPath)
	if err != nil {
		panic(err)
	}
}
