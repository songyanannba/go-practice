package business

import (
	uuid "github.com/satori/go.uuid"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/service/cache"
	"slot-server/utils"
)

type UserService struct {
}

// CreateUser 创建User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user business.User) (err error) {
	user.Password = utils.BcryptHash(user.Password)
	user.Uuid = uuid.NewV4().String()
	err = global.GVA_DB.Create(&user).Error
	return err
}

// DeleteUser 删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUser(user business.User) (err error) {
	err = global.GVA_DB.Delete(&user).Error
	return err
}

// DeleteUserByIds 批量删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.User{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUser 更新User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) UpdateUser(user business.User) (err error) {
	err = global.GVA_DB.Omit("uuid", "password").Save(&user).Error
	return err
}

func (userService *UserService) ChangePassword(user business.User) (err error) {
	if err = global.GVA_DB.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return err
	}
	user.Password = utils.BcryptHash(user.Password)
	err = global.GVA_DB.Select("password").Save(&user).Error
	return err
}

func (userService *UserService) ChangeAmount(user business.User) (err error) {
	_, err = cache.ChangeMoney(user.ID, user.Amount, nil)
	return err
}

// GetUser 根据id获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUser(id uint) (user business.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&user).Error
	user.Password = "******"
	return
}

// GetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUserInfoList(info businessReq.UserSearch) (list []business.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.User{})
	var users []business.User
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%"+info.NickName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}
	if info.Amount > 0 {
		db = db.Where("amount > ?", info.Amount)
	}
	if info.Status != 0 {
		db = db.Where("enable = ?", info.Status)
	}
	if info.Online != 0 {
		db = db.Where("online = ?", info.Online)
	}
	if info.MerchantId != 0 {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["amount"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&users).Error
	for i, user := range users {
		user.Password = "******"
		users[i] = user
	}
	return users, total, err
}
