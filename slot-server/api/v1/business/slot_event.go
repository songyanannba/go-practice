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
    "slot-server/utils"
)

type SlotEventApi struct {
}

var slotEventService = service.ServiceGroupApp.BusinessServiceGroup.SlotEventService


// CreateSlotEvent 创建SlotEvent
// @Tags SlotEvent
// @Summary 创建SlotEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotEvent true "创建SlotEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotEvent/createSlotEvent [post]
func (slotEventApi *SlotEventApi) CreateSlotEvent(c *gin.Context) {
	var slotEvent business.SlotEvent
	err := c.ShouldBindJSON(&slotEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "SlotId":{utils.NotEmpty()},
    }
	if err := utils.Verify(slotEvent, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := slotEventService.CreateSlotEvent(slotEvent); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSlotEvent 删除SlotEvent
// @Tags SlotEvent
// @Summary 删除SlotEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotEvent true "删除SlotEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /slotEvent/deleteSlotEvent [delete]
func (slotEventApi *SlotEventApi) DeleteSlotEvent(c *gin.Context) {
	var slotEvent business.SlotEvent
	err := c.ShouldBindJSON(&slotEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotEventService.DeleteSlotEvent(slotEvent); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSlotEventByIds 批量删除SlotEvent
// @Tags SlotEvent
// @Summary 批量删除SlotEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SlotEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /slotEvent/deleteSlotEventByIds [delete]
func (slotEventApi *SlotEventApi) DeleteSlotEventByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := slotEventService.DeleteSlotEventByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSlotEvent 更新SlotEvent
// @Tags SlotEvent
// @Summary 更新SlotEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body business.SlotEvent true "更新SlotEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /slotEvent/updateSlotEvent [put]
func (slotEventApi *SlotEventApi) UpdateSlotEvent(c *gin.Context) {
	var slotEvent business.SlotEvent
	err := c.ShouldBindJSON(&slotEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "SlotId":{utils.NotEmpty()},
      }
    if err := utils.Verify(slotEvent, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := slotEventService.UpdateSlotEvent(slotEvent); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败: " + err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSlotEvent 用id查询SlotEvent
// @Tags SlotEvent
// @Summary 用id查询SlotEvent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query business.SlotEvent true "用id查询SlotEvent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /slotEvent/findSlotEvent [get]
func (slotEventApi *SlotEventApi) FindSlotEvent(c *gin.Context) {
	var slotEvent business.SlotEvent
	err := c.ShouldBindQuery(&slotEvent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reslotEvent, err := slotEventService.GetSlotEvent(slotEvent.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败: " + err.Error(), c)
	} else {
		response.OkWithData(gin.H{"reslotEvent": reslotEvent}, c)
	}
}

// GetSlotEventList 分页获取SlotEvent列表
// @Tags SlotEvent
// @Summary 分页获取SlotEvent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query businessReq.SlotEventSearch true "分页获取SlotEvent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /slotEvent/getSlotEventList [get]
func (slotEventApi *SlotEventApi) GetSlotEventList(c *gin.Context) {
	var pageInfo businessReq.SlotEventSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := slotEventService.GetSlotEventInfoList(pageInfo); err != nil {
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
