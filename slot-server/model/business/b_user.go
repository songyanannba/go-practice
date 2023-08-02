// 自动生成模板User
package business

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"slot-server/global"
)

// User 结构体
type User struct {
	global.GVA_MODEL
	Username   string `json:"username" form:"username" gorm:"index;column:username;comment:用户名;size:30;"`
	Password   string `json:"password" form:"password" gorm:"column:password;comment:密码;size:200;"`
	Uuid       string `json:"uuid" form:"uuid" gorm:"index;column:uuid;comment:UUID;"`
	NickName   string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:昵称;size:50;"`
	Phone      string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:30;"`
	Email      string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:100;"`
	HeaderImg  string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:头像;size:255;"`
	Status     uint8  `json:"status" form:"status" gorm:"column:status;default:1;comment:状态 1正常 2冻结;size:8;"`
	MerchantId uint   `json:"merchantId" form:"merchantId" gorm:"column:merchant_id;default:0;comment:商户ID;size:32;"`

	Ip       string `json:"ip" form:"ip" gorm:"column:ip;comment:注册IP;size:30;"`
	LastIp   string `json:"lastIp" form:"lastIp" gorm:"column:last_ip;comment:最后登录IP;size:30;"`
	Amount   int64  `json:"amount" form:"amount" gorm:"column:amount;default:0;comment:金额;size:64;"`
	Online   uint8  `json:"online" form:"online" gorm:"column:online;default:2;comment:是否在线;size:8;"`
	Currency string `json:"currency" form:"currency" gorm:"column:currency;default:USD;comment:货币;size:12;"`

	Token string `json:"token" form:"token" gorm:"-"`
}

// TableName User 表名
func (User) TableName() string {
	return "b_user"
}

func (u *User) ValidateSignUp() error {
	return validation.ValidateStruct(u) //validation.Field(&u.Username, validation.Required, validation.Length(5, 20)),
	//validation.Field(&u.Password, validation.Required, validation.Length(5, 20)),
	//validation.Field(&u.NickName, validation.Required, validation.Length(5, 30)),
	//validation.Field(&u.Phone, validation.Required, validation.Length(5, 20)),

}

func (u *User) ValidateLogin() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required, validation.Length(5, 20)),
		//validation.Field(&u.Password, validation.Required, validation.Length(5, 20)),
	)
}

func (u *User) GetCurrency() {
	if u.ID == 0 {
		u.Currency = "USD"
		return
	}
	var currency string
	global.GVA_DB.Model(&User{}).Where("id = ?", u.ID).Pluck("currency", &currency)
	if currency == "" {
		u.Currency = "USD"
		return
	}
	u.Currency = currency
	return
}
