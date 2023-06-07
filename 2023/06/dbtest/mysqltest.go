package dbtest

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

// 测试mysql的读写
type DbUser struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	Age       int            `json:"age" form:"age" gorm:"column:age"`
	First     string         `json:"first" form:"first" gorm:"column:first"`
	Last      string         `json:"last" form:"last" gorm:"column:last"`
	CreatedAt time.Time      `gorm:"size:0"`  // 创建时间
	UpdatedAt time.Time      `gorm:"size:0"`  // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"size:0" ` // 删除时间
}

func (DbUser) TableName() string {
	return "user"
}

func WriteMsqUser() {

	du := DbUser{
		Age:       12,
		First:     "111",
		Last:      "222",
	}

	DB.Create(&du)
}


func Read(uid int) {

	var u1 DbUser
	if err = DB.Model(&DbUser{}).Where("id = ?", uid).First(&u1).Error; err != nil {
		fmt.Println(err)
	}

	//fmt.Println("u1 == " , u1)

	/*var us []DbUser
	if err = DB.Model(&DbUser{}).Where("id > ?", 1).Find(&us).Error; err != nil {
		fmt.Println(err)
	}*/

}
