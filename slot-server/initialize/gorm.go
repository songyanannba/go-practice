package initialize

import (
	"os"

	"slot-server/global"
	"slot-server/model/example"
	"slot-server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"slot-server/model/business"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		// 示例模块表
		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		// 自动化模块表
		// Code generated by slot-server Begin; DO NOT EDIT.

		business.User{},
		business.Slot{},
		business.SlotReel{},
		business.SlotSymbol{},
		business.SlotPayTable{},
		business.SlotPayline{},
		business.Jackpot{},
		business.SlotTests{},
		business.Configs{},
		business.SlotRecord{},
		business.MoneyLog{},
		business.SlotUserSpin{},
		business.SlotEvent{},
		business.MoneySlot{},
		business.MoneyUser{},
		business.SlotFake{},
		business.MoneySlot{},
		business.MoneyUserSlot{},
		business.SlotReelData{},
		business.Merchant{},
		business.Tracking{},
		business.ApiLog{},
		business.DebugConfig{},
		business.Txn{},
		business.SlotFileUploadAndDownload{},
		business.TxnSub{},
		business.Currency{},
		business.SlotGenTpl{},
		business.SlotTemplate{},
		business.SlotTemplateGen{},
		// Code generated by slot-server End; DO NOT EDIT.
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}