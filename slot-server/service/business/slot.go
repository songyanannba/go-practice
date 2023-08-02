package business

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"slot-server/config"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	businessReq "slot-server/model/business/request"
	"slot-server/model/common/request"
	"slot-server/model/common/response"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/cluster/backend"
	"slot-server/service/logic"
	slotGame "slot-server/service/slot/component"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strconv"
	"strings"
)

type SlotService struct {
}

// CreateSlot 创建Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService) CreateSlot(slot business.Slot) (err error) {
	err = global.GVA_DB.Create(&slot).Error
	slotGame.DeleteConfigCache(slot.ID)
	return err
}

// DeleteSlot 删除Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService) DeleteSlot(slot business.Slot) (err error) {
	err = global.GVA_DB.Delete(&slot).Error
	slotGame.DeleteConfigCache(slot.ID)
	return err
}

// DeleteSlotByIds 批量删除Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService) DeleteSlotByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]business.Slot{}, "id in ?", ids.Ids).Error
	for _, id := range ids.Ids {
		slotGame.DeleteConfigCache(uint(id))
	}
	return err
}

// UpdateSlot 更新Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService) UpdateSlot(slot business.Slot) (err error) {
	err = global.GVA_DB.Save(&slot).Error
	slotGame.DeleteConfigCache(slot.ID)
	return err
}

// GetSlot 根据id获取Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService) GetSlot(id uint) (slot business.Slot, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&slot).Error
	return
}

// GetSlotInfoList 分页获取Slot记录
// Author [piexlmax](https://github.com/piexlmax)
func (slotService *SlotService) GetSlotInfoList(info businessReq.SlotSearch) (list []business.Slot, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&business.Slot{})
	var slots []business.Slot
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.BetweenTime) > 1 {
		db = db.Where("created_at BETWEEN ? and ?", info.BetweenTime[0], info.BetweenTime[1])
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.PaylineNo != 0 {
		db = db.Where("payline_no = ?", info.PaylineNo)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
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
	}

	err = db.Limit(limit).Offset(offset).Find(&slots).Error
	return slots, total, err
}

type AllConfigParams struct {
	Id   int `json:"slot_id"`
	Data struct {
		Slot       []*business.Slot         `json:"slot"`
		Reel       *SlotReelConfig          `json:"slot_reel"`
		PayTable   []*business.SlotPayTable `json:"slot_pay_table"`
		Payline    []*business.SlotPayline  `json:"slot_payline"`
		SlotSymbol []*business.SlotSymbol   `json:"slot_symbol"`
		SlotEvent  []*business.SlotEvent    `json:"slot_event"`
		SlotFake   []*business.SlotFake     `json:"slot_fake"`
	} `json:"data"`
}

type SlotReelConfig struct {
	Reels     []*business.SlotReel     `json:"reels"`
	ReelDatas []*business.SlotReelData `json:"reel_datas"`
}

// UpdateAllConfig 更新所有配置
func (slotService *SlotService) UpdateAllConfig(params *AllConfigParams) (text string, success bool) {
	success = true
	l := log{}
	logs := make([]log, 0)
	defer func() {
		str := SlotEventVerify(params.Data.SlotEvent)
		if str != "" {
			success = false
			text = str
			return
		}

		var msg []string
		for _, l := range logs {
			msg = append(msg, l.GetText())
		}
		text = strings.Join(msg, "<br>")
		return
	}()

	l, success = handleUploadData(business.Slot{}.TableName(), params.Data.Slot)
	logs = append(logs, l)
	if !success {
		return
	}

	l, success = handleUploadData(business.SlotReel{}.TableName(), params.Data.Reel.Reels)
	logs = append(logs, l)
	if !success {
		return
	}

	l, success = handleUploadData(business.SlotReelData{}.TableName(), params.Data.Reel.ReelDatas)
	logs = append(logs, l)
	if !success {
		return
	}

	l, success = handleUploadData(business.SlotPayTable{}.TableName(), params.Data.PayTable)
	logs = append(logs, l)
	if !success {
		return
	}

	l, success = handleUploadData(business.SlotPayline{}.TableName(), params.Data.Payline)
	logs = append(logs, l)
	if !success {
		return
	}

	l, success = handleUploadData(business.SlotSymbol{}.TableName(), params.Data.SlotSymbol)
	logs = append(logs, l)
	if !success {
		return
	}

	l, success = handleUploadData(business.SlotEvent{}.TableName(), params.Data.SlotEvent)
	logs = append(logs, l)
	if !success {
		return
	}

	l, success = handleUploadData(business.SlotFake{}.TableName(), params.Data.SlotFake)
	logs = append(logs, l)
	if !success {
		return
	}

	// 删除本地所有缓存 (游戏服务器需后台发送清除请求)
	slotGame.FlushConfigCache()
	return
}

func SlotEventVerify(events []*business.SlotEvent) string {
	demoEvents := lo.Filter(events, func(item *business.SlotEvent, index int) bool {
		return item.Demo == 1
	})

	Events := lo.Filter(events, func(item *business.SlotEvent, index int) bool {
		return item.Demo == 0
	})
	for i, event := range demoEvents {
		if i == 52 {
			continue
		}
		strs := strings.Split(event.Event1, "@")
		if len(strs) != 2 {
			return fmt.Sprintf("DemoEvent第%d行事件1格式错误", i+1)
		}
		str1s := strings.Split(strs[0], "&")
		str2s := strings.Split(strs[1], "&")
		if len(str1s)+1 != len(str2s) {
			return fmt.Sprintf("DemoEvent第%d行事件1格式错误", i+1)
		}
	}

	for i, event := range Events {
		if i == 52 {
			continue
		}
		strs := strings.Split(event.Event1, "@")
		if len(strs) != 2 {
			return fmt.Sprintf("第%d行事件1格式错误", i+1)
		}
		str1s := strings.Split(strs[0], "&")
		str2s := strings.Split(strs[1], "&")
		if len(str1s)+1 != len(str2s) {
			return fmt.Sprintf("第%d行事件1格式错误", i+1)
		}
	}

	return ""
}

type log struct {
	TableName string
	Success   bool
	Msg       string
	Insert    int64
	Del       int64
}

func (l log) GetText() string {
	return fmt.Sprintf(
		`%s表:`+
			` %s 新增数据[%s]条 删除数据[%s]条 %s`,
		helper.GreenTag(l.TableName),
		helper.If(l.Success, helper.GreenTag("成功"), helper.RedTag("失败")),
		helper.OrangeTag(strconv.FormatInt(l.Insert, 10)),
		helper.OrangeTag(strconv.FormatInt(l.Del, 10)),
		l.Msg,
	)
}

func handleUploadData[T business.Table](table string, data []T) (l log, success bool) {
	success = true
	l = log{
		TableName: table[2:],
		Success:   true,
	}
	if len(data) == 0 {
		return
	}
	global.GVA_DB.Raw("select count(*) from `" + table + "`").Scan(&l.Del)
	global.GVA_DB.Exec("truncate table `" + table + "`")
	err := global.GVA_DB.CreateInBatches(data, 1000).Error
	if err != nil {
		l.Success = false
		success = false
		l.Msg = "error : " + err.Error()
	}
	l.Insert = int64(len(data))
	return
}

type BackendOperateParams struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

func (slotService *SlotService) BackendOperate(token string, params *BackendOperateParams) (string, error) {
	var (
		msgs []string
	)
	switch params.Type {
	case enum.BackendOperateType1RefreshCache:
		cache.ClearLocalCache() // 清除本地缓存
		//business.DelAllXConfigCache() // 清除xconfig缓存
		msgs = append(msgs, reqLog(reqLogParams{Name: "backend", Url: "127.0.0.1"}))
	}

	for _, cluster := range global.GVA_CONFIG.System.Clusters {
		msgs = append(msgs, sendReq(token, cluster, params))
	}
	return strings.Join(msgs, "<br>"), nil
}

// 发送后台操作请求至所有集群
func sendReq(token string, cluster config.Cluster, params *BackendOperateParams) (res string) {
	var (
		gateLog   = reqLogParams{Name: cluster.Name + "-gate", Url: cluster.GetUrl(global.GVA_CONFIG.System.GateAddr)}
		gameLog   = reqLogParams{Name: cluster.Name + "-game", Url: cluster.GetUrl(global.GVA_CONFIG.System.GameAddr)}
		apiLog    = reqLogParams{Name: cluster.Name + "-api", Url: cluster.GetUrl(global.GVA_CONFIG.System.ApiAddr)}
		commonMsg string
	)
	defer func() {
		if commonMsg != "" {
			gateLog.Msg = commonMsg
			gameLog.Msg = commonMsg
			apiLog.Msg = commonMsg
		}
		res = reqLog(gateLog, gameLog, apiLog)
		return
	}()
	err := apiReq("http://"+apiLog.Url+"/game/backendOperate", params)
	if err != nil {
		commonMsg = err.Error()
		return
	}

	conn, ok := backend.ClusterConnPool[cluster.Name]
	if !ok {
		gateLog.Msg = "集群连接不存在"
		gameLog.Msg = "集群连接不存在"
		return
	}
	ack, err := logic.ReqBackendOperate(conn, token, params.Type, params.Data)
	if err != nil {
		gateLog.Msg = err.Error()
		gameLog.Msg = err.Error()
		return
	}
	if ack.Head.Code != pbs.Code_Ok {
		if ack.Head.Code == pbs.Code_Unknown {
			gameLog.Msg = ack.Head.Message
		} else {
			gateLog.Msg = ack.Head.Message
			gameLog.Msg = ack.Head.Message
		}
	}
	return
}

func apiReq(url string, params *BackendOperateParams) error {
	jsonData, _ := global.Json.Marshal(params.Data)
	req := &pbs.BackendOperate{
		Head: &pbs.ReqHead{Token: enum.AdminDefaultToken},
		Type: int32(params.Type),
		Data: string(jsonData),
	}
	apiRes, err := utils.NewGurl("POST", url).SetData(req).Do()
	if err != nil {
		return err
	}
	var res response.Response
	err = global.Json.Unmarshal(apiRes, &res)
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return errors.New(res.Msg)
	}
	return nil
}

type reqLogParams struct {
	Name string
	Url  string
	Msg  string
}

func reqLog(params ...reqLogParams) string {
	var (
		logs []string
	)
	for _, param := range params {
		logs = append(logs, fmt.Sprintf(
			"[%s] %s 状态 %s %s",
			param.Name,
			helper.GreenTag(param.Url),
			helper.If(param.Msg == "", helper.GreenTag("成功"), helper.RedTag("失败")),
			helper.OrangeTag(param.Msg),
		))
	}
	return strings.Join(logs, "<br>")
}
