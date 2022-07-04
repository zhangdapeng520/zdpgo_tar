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

	// compress
	err := tar.TarGzDirAllFiles("examples/tar_dir")
	if err != nil {
		panic(err)
	}

	// uncompress
	err = tar.UnTarGzDir("examples/tar_dir")
	if err != nil {
		panic(err)
	}
}
