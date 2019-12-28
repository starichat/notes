package main

import (
	"flag"
	"fmt"
    "github.com/spf13/viper"
)

var (
	config = flag.String("config", "config", "配置文件名称，默认 config")
)

func main() {
	// 配置文件名称
	viper.SetConfigName(*config)
	// 查找文件,可以配置多个路径
	viper.AddConfigPath("./")
	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: % \n",err))
	}
	// 监控文件变化
	viper.WatchConfig()

	environment := viper.GetBool("security.enabled")
	fmt.Println("security.enabled: ",environment)


}