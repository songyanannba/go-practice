package business

import (
	"mime/multipart"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/utils/upload"
	"strconv"
	"strings"
)

type SlotFileUploadAndDownloadService struct {
}

// CreateSlotFileUploadAndDownload 创建SlotFileUploadAndDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlotFileUpAndDownService *SlotFileUploadAndDownloadService) CreateSlotFileUploadAndDownload(SlotFileUpAndDown business.SlotFileUploadAndDownload) (err error) {
	err = global.GVA_DB.Create(&SlotFileUpAndDown).Error
	return err
}

// DeleteSlotFileUploadAndDownload 删除SlotFileUploadAndDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlotFileUpAndDownService *SlotFileUploadAndDownloadService)DeleteSlotFileUploadAndDownload(SlotFileUpAndDown business.SlotFileUploadAndDownload) (err error) {
	err = global.GVA_DB.Delete(&SlotFileUpAndDown).Error
	return err
}

// DeleteSlotFileUploadAndDownloadByIds 批量删除SlotFileUploadAndDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlotFileUpAndDownService *SlotFileUploadAndDownloadService)DeleteSlotFileUploadAndDownloadByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.SlotFileUploadAndDownload{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSlotFileUploadAndDownload 更新SlotFileUploadAndDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlotFileUpAndDownService *SlotFileUploadAndDownloadService)UpdateSlotFileUploadAndDownload(SlotFileUpAndDown business.SlotFileUploadAndDownload) (err error) {
	err = global.GVA_DB.Save(&SlotFileUpAndDown).Error
	return err
}

// GetSlotFileUploadAndDownload 根据id获取SlotFileUploadAndDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlotFileUpAndDownService *SlotFileUploadAndDownloadService)GetSlotFileUploadAndDownload(id uint) (SlotFileUpAndDown business.SlotFileUploadAndDownload, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&SlotFileUpAndDown).Error
	return
}

// GetSlotFileUploadAndDownloadInfoList 分页获取SlotFileUploadAndDownload记录
// Author [piexlmax](https://github.com/piexlmax)
func (SlotFileUpAndDownService *SlotFileUploadAndDownloadService)GetSlotFileUploadAndDownloadInfoList(info businessReq.SlotFileUploadAndDownloadSearch) (list []business.SlotFileUploadAndDownload, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&business.SlotFileUploadAndDownload{})
    var SlotFileUpAndDowns []business.SlotFileUploadAndDownload
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.BetweenTime) > 1 {
    	db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
	if info.FileDir != "" {
		db = db.Where("file_dir = ?",info.FileDir)
	}
    if info.Type != 0 {
        db = db.Where("type = ?",info.Type)
    }
    if info.SlotId != 0 {
        db = db.Where("slot_id = ?",info.SlotId)
    }
	if info.Specification != "" {
		db = db.Where("specification = ?", info.Specification)
	}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
    var OrderStr string
    orderMap := make(map[string]bool)
    orderMap["id"] = true
    if orderMap[info.Sort] {
       OrderStr = "`" + info.Sort + "`"
       if info.Order == "descending" {
          OrderStr = OrderStr + " desc"
       }
       db = db.Order(OrderStr)
    }  else {
      	db = db.Order("`id` desc")
    }

	err = db.Limit(limit).Offset(offset).Find(&SlotFileUpAndDowns).Error
	return  SlotFileUpAndDowns, total, err
}


func (e *SlotFileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (file business.SlotFileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := business.SlotFileUploadAndDownload{
			Url:  filePath,
			Name: s[0],
			//Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return f, e.CreateSlotFileUploadAndDownload(f)
	}
	return
}

func (e *SlotFileUploadAndDownloadService) UploadFileDir(header *multipart.FileHeader, noSave, fileDir string, userId int) (file business.SlotFileUploadAndDownload, err error) {
	filePath, key, uploadErr := upload.UploadFileDir(header, fileDir)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		sId, _ := strconv.Atoi(fileDir)
		f := business.SlotFileUploadAndDownload{
			Url:  filePath,
			Name: s[0],
			//Name: header.Filename,
			Tag:     s[len(s)-1],
			Key:     key,
			FileDir: fileDir,
			SlotId:  sId,
			UserId:  userId,
		}
		return e.CreateOrUpdatedb(f)
		//return f, e.CreateSlotFileUploadAndDownload(f)
	}
	return
}

func (e *SlotFileUploadAndDownloadService) CreateOrUpdatedb(f business.SlotFileUploadAndDownload) (file business.SlotFileUploadAndDownload, err error) {
	SlotFileUpAndDown := business.SlotFileUploadAndDownload{}
	//先判断是不是修改
	err = global.GVA_DB.Where("name = ? ", f.Name).Where("slot_id = ? ", f.SlotId).Where("file_dir = ? ", f.FileDir).First(&SlotFileUpAndDown).Error
	if SlotFileUpAndDown.ID > 0 {
		e.UpdateSlotFileUploadAndDownload(SlotFileUpAndDown)
		return SlotFileUpAndDown , err
	}
	return f, e.CreateSlotFileUploadAndDownload(f)
}
