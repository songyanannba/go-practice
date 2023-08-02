package core

import (
	"fmt"
	"slot-server/config"
	"strings"
	"time"

	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/initialize"
	"slot-server/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info(fmt.Sprintf(
		"server run success on %s,\n"+
			"默认自动化文档地址:http://127.0.0.1%s/swagger/index.html\n"+
			"默认前端文件运行地址:http://127.0.0.1:8080",
		address, address,
	))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

func RunApiServer() {
	Router := initialize.ApiRouters()

	address := global.GVA_CONFIG.System.ApiAddr
	if !strings.Contains(address, ":") {
		address = ":" + address
	}
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info(fmt.Sprintf(
		"server run success on %s,\n"+
			"api地址:http://127.0.0.1%s,\n"+
			"域名:%s\n",
		address, address, global.GVA_CONFIG.System.ApiDomain,
	))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

func BaseInit() {
	global.GVA_VP = Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG.Logger)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
}

func TestModInit() {
	global.GVA_CONFIG = config.Server{
		Redis:  config.Redis{},
		System: config.System{},
		Mysql: config.Mysql{
			GeneralDB: config.GeneralDB{
				Path:         "47.99.106.90",
				Port:         "3306",
				Config:       "charset=utf8mb4&parseTime=True&loc=Local",
				Dbname:       "slot",
				Username:     "root",
				Password:     "Huang1998",
				Prefix:       "",
				Singular:     false,
				Engine:       "",
				MaxIdleConns: 10,
				MaxOpenConns: 100,
				LogMode:      "info",
				LogZap:       false,
			},
		},
	}
	initialize.OtherInit()
	global.GVA_LOG = Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG.Logger)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
}
