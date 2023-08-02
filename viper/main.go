package main

import (
	"fmt"

	"github.com/qppHUST/prepareForOffer/viper/model"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("D:\\workspace\\vscode_demo\\prepareForOffer\\viper\\config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := model.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
}
