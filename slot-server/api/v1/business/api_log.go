package business

import (
	"slot-server/global"
    "slot-server/model/business"
    "slot-server/model/common/request"
    businessReq "slot-server/model/business/request"
    "slot-server/model/common/response"
    "slot-server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ApiLogApi struct {
}

var apiLogService = service.ServiceGroupApp.BusinessServiceGroup.ApiLogService


// CreateApiLog 创建ApiLog
// @Tags ApiLog
// @Summary 创建ApiLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ApiLog true "创建ApiLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiLog/createApiLog [post]
func (apiLogApi *ApiLogApi) CreateApiLog(c *gin.Context) {
	var apiLog business.ApiLog
	err := c.ShouldBindJSON(&apiLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiLogService.CreateApiLog(apiLog); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApiLog 删除ApiLog
// @Tags ApiLog
// @Summary 删除ApiLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ApiLog true "删除ApiLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiLog/deleteApiLog [delete]
func (apiLogApi *ApiLogApi) DeleteApiLog(c *gin.Context) {
	var apiLog business.ApiLog
	err := c.ShouldBindJSON(&apiLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiLogService.DeleteApiLog(apiLog); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApiLogByIds 批量删除ApiLog
// @Tags ApiLog
// @Summary 批量删除ApiLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apiLog/deleteApiLogByIds [delete]
func (apiLogApi *ApiLogApi) DeleteApiLogByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiLogService.DeleteApiLogByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApiLog 更新ApiLog
// @Tags ApiLog
// @Summary 更新ApiLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.ApiLog true "更新ApiLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiLog/updateApiLog [put]
func (apiLogApi *ApiLogApi) UpdateApiLog(c *gin.Context) {
	var apiLog business.ApiLog
	err := c.ShouldBindJSON(&apiLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := apiLogService.UpdateApiLog(apiLog); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApiLog 用id查询ApiLog
// @Tags ApiLog
// @Summary 用id查询ApiLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.ApiLog true "用id查询ApiLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiLog/findApiLog [get]
func (apiLogApi *ApiLogApi) FindApiLog(c *gin.Context) {
	var apiLog business.ApiLog
	err := c.ShouldBindQuery(&apiLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reapiLog, err := apiLogService.GetApiLog(apiLog.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reapiLog": reapiLog}, c)
	}
}

// GetApiLogList 分页获取ApiLog列表
// @Tags ApiLog
// @Summary 分页获取ApiLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.ApiLogSearch true "分页获取ApiLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiLog/getApiLogList [get]
func (apiLogApi *ApiLogApi) GetApiLogList(c *gin.Context) {
	var pageInfo businessReq.ApiLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := apiLogService.GetApiLogInfoList(pageInfo); err != nil {
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
