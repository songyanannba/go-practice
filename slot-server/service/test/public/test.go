package public

import (
	"fmt"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/utils/helper"
)

type RunSlotTest struct {
	Type       int    `json:"type"`
	SlotId     uint   `json:"slotId"`
	Num        int    `json:"num"`
	Hold       int    `json:"hold"` //持有金币
	Amount     int    `json:"amount"`
	Opts       []int  `json:"opts"`
	Raise      int    `json:"raise"`
	Result     string `json:"result"`
	IsMustFree int    `json:"isMustFree"` // 购买免费次数
	IsMustRes  int    `json:"isMustRes"`  // 购买Respin次数
}

const SlotTestLock = "slot_test_lock"

type BasicData struct {
	N0  float64 //最高倍率:::::
	N1  int     //总赢钱::::
	N2  int     //总押注消耗钱:::::
	N3  int     //总转动次数:::::
	N4  int     //总赢钱次数:::::
	N5  int     //普通转划线赢钱1-5倍的次数:::::
	N6  int     //普通转划线赢钱1-5倍的总钱数:::::
	N7  int     //普通转划线赢钱5-10倍的次数:::::
	N8  int     //普通转划线赢钱5-10倍的总钱数:::::
	N9  int     //普通转划线赢钱10-20倍的次数:::::
	N10 int     //普通转划线赢钱10-20倍的总钱数:::::
	N11 int     //普通转划线赢钱20-30倍的次数:::::
	N12 int     //普通转划线赢钱20-30倍的总钱数:::::
	N13 int     //普通转划线赢钱30-40倍的次数:::::
	N14 int     //普通转划线赢钱30-40倍的总钱数:::::
	N15 int     //普通转划线赢钱40-50倍的次数:::::
	N16 int     //普通转划线赢钱40-50倍的总钱数:::::
	N17 int     //普通转划线赢钱50-100倍的次数:::::
	N18 int     //普通转划线赢钱50-100倍的总钱数:::::
	N19 int     //普通转划线赢钱100-200倍的次数:::::
	N20 int     //普通转划线赢钱100-200倍的总钱数:::::
	N21 int     //普通转划线赢钱200-500倍的次数:::::
	N22 int     //普通转划线赢钱200-500倍的总钱数:::::
	N23 int     //普通转划线赢钱500-1000倍的次数:::::
	N24 int     //普通转划线赢钱500-1000倍的总钱数:::::
	N25 int     //普通转划线赢钱1000以上倍的次数:::::
	N26 int     //普通转划线赢钱1000以上倍的总钱数:::::
}

// EndProcessing 结束处理
func EndProcessing(slotTest *business.SlotTests, testSlotResult *BasicData) {
	defer func() {
		global.BlackCache.Delete(SlotTestLock)
		if err := recover(); err != nil {
			global.GVA_LOG.Error("测试惊恐:" + fmt.Sprintf("%v", err))
			e := fmt.Sprintf("%v", err)
			fmt.Println(helper.Stack())
			slotTest.Detail = e
			slotTest.Status = enum.CommonStatusError
			global.GVA_DB.Save(&slotTest)
			return
		}
	}()
	slotTest.Win = testSlotResult.N1
	slotTest.RunNum = testSlotResult.N3
	slotTest.Status = enum.CommonStatusFinish
	err := global.GVA_DB.Save(&slotTest).Error
	if err != nil {
		global.GVA_DB.Model(&slotTest).Where("id = ?", slotTest.ID).Update("Status", enum.CommonStatusError)
		global.GVA_LOG.Error("测试错误:" + err.Error())
	}
}
