package initializer

import (
	"SH-admin/global"
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("加載config檔失敗：%s", err))
	}
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("載入至結構體失敗：%s", err))
	}
}
