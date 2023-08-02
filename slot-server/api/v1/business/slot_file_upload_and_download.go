package business

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/model/common/response"
	"slot-server/service"
	"slot-server/utils"
)

type SlotFileUploadAndDownloadApi struct {
}

var SlotFileUpAndDownService = service.ServiceGroupApp.BusinessServiceGroup.SlotFileUploadAndDownloadService


// CreateSlotFileUploadAndDownload 创建SlotFileUploadAndDownload
// @Tags SlotFileUploadAndDownload
// @Summary 创建SlotFileUploadAndDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotFileUploadAndDownload true "创建SlotFileUploadAndDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SlotFileUpAndDown/createSlotFileUploadAndDownload [post]
func (SlotFileUpAndDownApi *SlotFileUploadAndDownloadApi) CreateSlotFileUploadAndDownload(c *gin.Context) {
	var SlotFileUpAndDown business.SlotFileUploadAndDownload
	err := c.ShouldBindJSON(&SlotFileUpAndDown)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SlotFileUpAndDownService.CreateSlotFileUploadAndDownload(SlotFileUpAndDown); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotFileUploadAndDownload 删除SlotFileUploadAndDownload
// @Tags SlotFileUploadAndDownload
// @Summary 删除SlotFileUploadAndDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotFileUploadAndDownload true "删除SlotFileUploadAndDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SlotFileUpAndDown/deleteSlotFileUploadAndDownload [delete]
func (SlotFileUpAndDownApi *SlotFileUploadAndDownloadApi) DeleteSlotFileUploadAndDownload(c *gin.Context) {
	var SlotFileUpAndDown business.SlotFileUploadAndDownload
	err := c.ShouldBindJSON(&SlotFileUpAndDown)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SlotFileUpAndDownService.DeleteSlotFileUploadAndDownload(SlotFileUpAndDown); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotFileUploadAndDownloadByIds 批量删除SlotFileUploadAndDownload
// @Tags SlotFileUploadAndDownload
// @Summary 批量删除SlotFileUploadAndDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotFileUploadAndDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /SlotFileUpAndDown/deleteSlotFileUploadAndDownloadByIds [delete]
func (SlotFileUpAndDownApi *SlotFileUploadAndDownloadApi) DeleteSlotFileUploadAndDownloadByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SlotFileUpAndDownService.DeleteSlotFileUploadAndDownloadByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotFileUploadAndDownload 更新SlotFileUploadAndDownload
// @Tags SlotFileUploadAndDownload
// @Summary 更新SlotFileUploadAndDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotFileUploadAndDownload true "更新SlotFileUploadAndDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SlotFileUpAndDown/updateSlotFileUploadAndDownload [put]
func (SlotFileUpAndDownApi *SlotFileUploadAndDownloadApi) UpdateSlotFileUploadAndDownload(c *gin.Context) {
	var SlotFileUpAndDown business.SlotFileUploadAndDownload
	err := c.ShouldBindJSON(&SlotFileUpAndDown)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := SlotFileUpAndDownService.UpdateSlotFileUploadAndDownload(SlotFileUpAndDown); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotFileUploadAndDownload 用id查询SlotFileUploadAndDownload
// @Tags SlotFileUploadAndDownload
// @Summary 用id查询SlotFileUploadAndDownload
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotFileUploadAndDownload true "用id查询SlotFileUploadAndDownload"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SlotFileUpAndDown/findSlotFileUploadAndDownload [get]
func (SlotFileUpAndDownApi *SlotFileUploadAndDownloadApi) FindSlotFileUploadAndDownload(c *gin.Context) {
	var SlotFileUpAndDown business.SlotFileUploadAndDownload
	err := c.ShouldBindQuery(&SlotFileUpAndDown)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reSlotFileUpAndDown, err := SlotFileUpAndDownService.GetSlotFileUploadAndDownload(SlotFileUpAndDown.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reSlotFileUpAndDown": reSlotFileUpAndDown}, c)
	}
}

// GetSlotFileUploadAndDownloadList 分页获取SlotFileUploadAndDownload列表
// @Tags SlotFileUploadAndDownload
// @Summary 分页获取SlotFileUploadAndDownload列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotFileUploadAndDownloadSearch true "分页获取SlotFileUploadAndDownload列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SlotFileUpAndDown/getSlotFileUploadAndDownloadList [get]
func (SlotFileUpAndDownApi *SlotFileUploadAndDownloadApi) GetSlotFileUploadAndDownloadList(c *gin.Context) {
	var pageInfo businessReq.SlotFileUploadAndDownloadSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := SlotFileUpAndDownService.GetSlotFileUploadAndDownloadInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败: " + err.Error(), c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

// UploadFile
// @Tags      ExaSlotFileUploadAndDownload
// @Summary   上传文件示例
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /slotFileUploadAndDownload/upload [post]
func (slotFileUploadAndDownloadApi *SlotFileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file business.SlotFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")

	fileDir := c.Request.PostForm.Get("file_dir")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}

	//用户ID
	claims, err := utils.GetClaims(c)
	userId := 0
	if claims.ID != 0 && err == nil  {
		userId = int(claims.ID)
	}

	file, err = SlotFileUpAndDownService.UploadFileDir(header, noSave, fileDir, userId) // 文件上传后拿到文件路径
	//file, err = SlotFileUpAndDownService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	type fileResponse struct {
		File business.SlotFileUploadAndDownload `json:"file"`
	}
	response.OkWithDetailed(fileResponse{File: file}, "上传成功", c)
}

