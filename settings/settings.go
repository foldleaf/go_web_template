package settings

import (
	"fmt"

	// "github.com/gin-gonic/gin"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	// viper.SetConfigFile("config.yaml")	//指定文件加后缀
	viper.SetConfigName("config") // 指定配置文件名称，不需要带后缀，会自动识别指定目录下相同的文件名
	viper.SetConfigType("yaml")   //指定配置文件类型，用于远程获取配置，本地时不生效
	viper.AddConfigPath(".")      //指定查找配置文件的路径（这里用相对路径）
	err = viper.ReadInConfig()    //读取文件配置
	if err != nil {
		// 读取配置信息失败
		fmt.Println("viper.ReadInConfig() 读取配置信息失败:", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改")
	})

	return

}
