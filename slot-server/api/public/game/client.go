package game

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/response"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/logic"
	"slot-server/utils"
	"strconv"
	"time"
)

func Ping(c *gin.Context) {
	c.Data(200, "text/plain", []byte(""))
}

func Time(c *gin.Context) {
	now := time.Now()
	l := now.Location().String()
	s := l + " " + now.Format("2006-01-02 15:04:05") + "\n"
	c.Data(200, "text/plain", []byte(s))
}

func Log(c *gin.Context) {
	params := map[string]string{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		return
	}
	if len(params) == 0 {
		return
	}
	key := "37u/3zffV7ufG5/fn1/f37+7nw=="
	for sign, text := range params {
		v := utils.MD5V([]byte(text + key))
		if v != sign {
			return
		}
		global.GVA_LOG.Info("client debug : " + text)
	}
}

func BackendOperate(c *gin.Context) {
	var (
		params pbs.BackendOperate
	)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = logic.ParseBackendToken(params.Head)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = logic.HandleBackendOperate(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func GetSlotRecord(c *gin.Context) {
	var (
		search businessReq.SlotRecordPublicSearch
		record business.SlotRecord
	)
	err := c.ShouldBindJSON(&search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	id := business.ParseSlotRecordId(search.No)
	if id == 0 {
		response.FailWithMessage("record not found", c)
		return
	}
	err = global.GVA_DB.Where("id = ?", id).First(&record).Error
	if err != nil {
		response.FailWithMessage("record not found", c)
		return
	}

	slot, err := cache.GetSlot(record.SlotId)
	if err != nil {
		response.FailWithMessage("record not found", c)
		return
	}

	response.OkWithData(map[string]any{
		"gameId":   strconv.Itoa(int(record.SlotId)),
		"name":     slot.NamePkg,
		"ack":      base64.StdEncoding.EncodeToString(record.Ack),
		"currency": record.Currency,
		"time":     record.CreatedAt.Unix(),
		"config":   slot.ClientConf,
	}, c)
}
