// 自动生成模板SlotFileUploadAndDownload
package business

import (
	"slot-server/global"
	
)

// SlotFileUploadAndDownload 结构体
type SlotFileUploadAndDownload struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:255;"`
      FileDir  string `json:"file_dir" form:"file_dir" gorm:"column:file_dir;comment:文件目录名;size:255;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址;"`
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文件标签;"`
      Key  string `json:"key" form:"key" gorm:"column:key;comment:文件编号;"`
      Type  int `json:"type" form:"type" gorm:"column:type;default:0;comment:图片类型;size:32;"`
      SlotId  int `json:"slotId" form:"slotId" gorm:"column:slot_id;default:0;comment:机器id;size:32;"`
      UserId  int `json:"userId" form:"userId" gorm:"column:user_id;default:0;comment:机器id;size:32;"`
      Specification  string `json:"specification" form:"specification" gorm:"column:specification;default:0;comment:规格;"`
}


// TableName SlotFileUploadAndDownload 表名
func (SlotFileUploadAndDownload) TableName() string {
  return "slot_file_upload_and_download"
}

