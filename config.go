package zdpgo_tar

/*
@Time : 2022/6/2 16:03
@Author : 张大鹏
@File : config.go
@Software: Goland2021.3.1
@Description:
*/

type Config struct {
	Debug       bool   `yaml:"debug" json:"debug"`
	LogFilePath string `yaml:"log_file_path" json:"log_file_path"`
}
