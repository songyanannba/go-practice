package timedtask

import (
	"go.uber.org/zap"
	"slot-server/global"
)

func TimedTask() {
	// 金币流水任务
	//MoneyLogTask()
	// 机台任务
	//SlotTask()
	// 交易任务
	TxnTask()
}

func MoneyLogTask() {
	//机台每日汇总
	if _, err := global.GVA_Timer.AddTaskByFunc("MoneySlotCalToday", "30 * * * *", MoneySlotCalToday); err != nil {
		global.GVA_LOG.Error("add timer Flow_Summary error:", zap.Error(err))
	}
	if _, err := global.GVA_Timer.AddTaskByFunc("MoneySlotCalYesterday", "5 0 * * *", MoneySlotCalYesterday); err != nil {
		global.GVA_LOG.Error("add timer MoneySlotCalYesterday error:", zap.Error(err))
	}
	//用户每日汇总
	if _, err := global.GVA_Timer.AddTaskByFunc("MoneyUserCalToday", "30 * * * *", MoneyUserCalToday); err != nil {
		global.GVA_LOG.Error("add timer MoneyUserCalToday error:", zap.Error(err))
	}
	if _, err := global.GVA_Timer.AddTaskByFunc("MoneyUserCalYesterday", "5 0 * * *", MoneyUserCalYesterday); err != nil {
		global.GVA_LOG.Error("add timer MoneyUserCalYesterday error:", zap.Error(err))
	}
	//机台用户每日汇总
	if _, err := global.GVA_Timer.AddTaskByFunc("MoneyUserSlotCalToday", "30 * * * *", MoneyUserSlotCalToday); err != nil {
		global.GVA_LOG.Error("add timer MoneyUserSlotCalToday error:", zap.Error(err))
	}
	if _, err := global.GVA_Timer.AddTaskByFunc("MoneyUserSlotCalYesterday", "5 0 * * *", MoneyUserSlotCalYesterday); err != nil {
		global.GVA_LOG.Error("add timer MoneyUserSlotCalYesterday error:", zap.Error(err))
	}
}

func SlotTask() {
	if _, err := global.GVA_Timer.AddTaskByFunc("FinishSpin", "* * * * *", FinishSpin); err != nil {
		global.GVA_LOG.Error("add timer FinishSpin error:", zap.Error(err))
	}
}

func TxnTask() {
	if _, err := global.GVA_Timer.AddTaskByFunc("FinishTxn", "*/3 * * * *", FinishTxn); err != nil {
		global.GVA_LOG.Error("add timer FinishTxn error:", zap.Error(err))
	}
}
