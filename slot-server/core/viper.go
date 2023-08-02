package core

import (
	"flag"
	"fmt"
	"slot-server/utils/env"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"slot-server/global"
	_ "slot-server/packfile"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	svType := flag.String("type", "gate", "the server type")
	flag.Parse()
	global.SvName = *svType

	config := env.GetConfigFileName()
	fmt.Printf("您正在使用%s环境,config的路径为%s\n", env.Mode, config)

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
