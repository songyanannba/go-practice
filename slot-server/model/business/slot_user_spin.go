// 自动生成模板SlotUserSpin
package business

import (
	"gorm.io/gorm"
	"slot-server/global"
	"strconv"
)

// SlotUserSpin 结构体
type SlotUserSpin struct {
	global.GVA_MODEL
	UserId  int64 `json:"userId" form:"userId" gorm:"index;column:user_id;default:0;comment:用户id;size:64;"`
	SlotId  uint  `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
	PlayNum uint  `redis:"playNum" json:"playNum" form:"playNum" gorm:"column:play_num;default:0;comment:已玩游戏次数;size:32;"`
	FreeNum uint  `redis:"freeNum" json:"freeNum" form:"freeNum" gorm:"column:free_num;default:0;comment:已免费玩次数;size:32;"`
}

// TableName SlotUserSpin 表名
func (SlotUserSpin) TableName() string {
	return "b_slot_user_spin"
}

func getUserSpinCacheKey(userId, slotId uint) string {
	return "slot_user_spin:" + strconv.Itoa(int(userId)) + ":" + strconv.Itoa(int(slotId))
}

func AddUserSpinNum(userId, slotId uint, free, freeNum, playNum int, tx *gorm.DB) error {
	data := map[string]interface{}{}
	if free != 0 {
		data["free"] = gorm.Expr("free + ?", free)
	}
	if playNum > 0 {
		data["play_num"] = gorm.Expr("play_num + ?", playNum)
	}
	if freeNum > 0 {
		data["free_num"] = gorm.Expr("free_num + ?", freeNum)
	}
	if len(data) == 0 {
		return nil
	}
	if tx == nil {
		tx = global.GVA_DB
	}
	res := tx.Model(&SlotUserSpin{}).Where("user_id = ? and slot_id = ?", userId, slotId).
		Updates(data)
	err := res.Error
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		if free < 0 {
			data["free"] = 0
		}
		data["user_id"] = userId
		data["slot_id"] = slotId
		data["created_at"] = gorm.Expr("NOW()")
		data["updated_at"] = gorm.Expr("NOW()")
		err = tx.Model(&SlotUserSpin{}).Create(data).Error
	}
	return err
}
