package seamlessWallet

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/url"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/model/common/response"
	"slot-server/service/cache"
	"slot-server/utils/helper"
	"sort"
	"strconv"
	"time"
)

type OpenHandle struct {
	Merchant *business.Merchant
	Params   AcceptHeadData
	Url      string
	Res      any
}

func NewOpenHandle(params AcceptHeadData, url string) *OpenHandle {
	return &OpenHandle{
		Params: params,
		Url:    url,
	}
}

func (o *OpenHandle) Check() error {
	head := o.Params.GetAcceptHead()
	err := head.CheckHead()
	if err != nil {
		return err
	}

	err = o.Params.CheckParams()
	if err != nil {
		return err
	}

	o.Merchant, err = head.CheckSign(o.Params)
	return err
}

type OpenGameParams struct {
	AcceptHead
	GameId     string `json:"gameId" example:"1"`                          // 游戏id
	Token      string `json:"token" example:"test_token"`                  // 由运营商提供的token，⽤于在Authenticate⾥ 进⾏验证
	Language   string `json:"language" example:"en"`                       // 语言 en
	Currency   string `json:"currency" example:"USD"`                      // 货币类型 USD
	LobbyUrl   string `json:"lobbyUrl" example:"https://www.google.com"`   // 返回按钮的对应连接
	CashierUrl string `json:"cashierUrl" example:"https://www.google.com"` // 破产之后弹出的充值⻚⾯
	Demo       bool   `json:"demo"`                                        // 是否为试玩模式
}

func (o *OpenGameParams) CheckParams() error {
	if o.LobbyUrl != "" {
		_, err := url.ParseRequestURI(o.LobbyUrl)
		if err != nil {
			return err
		}
	}
	if o.CashierUrl != "" {
		_, err := url.ParseRequestURI(o.CashierUrl)
		if err != nil {
			return err
		}
	}
	if err := CheckLanguage(o.Language); err != nil {
		return err
	}
	if len(o.Currency) > 10 {
		return errors.New("currency is too long")
	}
	return nil
}

func (o *OpenHandle) OpenGame() error {
	if err := o.Check(); err != nil {
		return err
	}
	params := o.Params.(*OpenGameParams)
	slotId, err := strconv.Atoi(params.GameId)
	if err != nil {
		return errors.New("game id does not exist")
	}

	slot, err := cache.GetSlot(uint(slotId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("gameId is not exist")
		}
		global.GVA_LOG.Error("get slot error: " + err.Error())
		return errors.New("get game id error")
	}

	if slot.Status != enum.Yes {
		return errors.New("the game is closing")
	}

	user, err := CheckMerchantToken(o.Merchant, params.Token)
	if err != nil {
		return err
	}

	url := global.GVA_CONFIG.System.GameDomain
	if url == "" {
		url = "https://h5.bigwin.money"
	}

	o.Res = url + "?game_id=" + params.GameId + "&token=" + params.Token +
		"&language=" + params.Language + "&currency=" + params.Currency +
		"&lobbyUrl=" + params.LobbyUrl + "&cashierUrl=" + params.CashierUrl +
		"&demo=" + strconv.FormatBool(params.Demo) + "&username=" + user.Username +
		"&platform=" + o.Merchant.Agent
	return nil
}

type GameListParams struct {
	AcceptHead
	Language string `json:"language" example:"en"` // 语言
}

func (o *GameListParams) CheckParams() error {
	if err := CheckLanguage(o.Language); err != nil {
		return err
	}
	return nil
}

type GameListRes struct {
	GameId        string `json:"gameId"`        // 游戏id
	Name          string `json:"name"`          // 游戏名称
	DemoAvailable bool   `json:"demoAvailable"` // 试玩模式是否可用
	Icon          string `json:"icon"`          // 游戏图标
}

func (o *OpenHandle) GameList() error {
	if err := o.Check(); err != nil {
		return err
	}
	slots, err := cache.GetSlotList()
	if err != nil {
		return err
	}
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].ID < slots[j].ID
	})
	var res []*GameListRes
	for _, slot := range slots {
		res = append(res, &GameListRes{
			GameId:        strconv.Itoa(int(slot.ID)),
			Name:          slot.Name,
			DemoAvailable: false,
			Icon:          global.GVA_CONFIG.System.StorageDomain + "/cover/icon_" + strconv.Itoa(int(slot.ID)) + ".png",
		})
	}
	o.Res = res
	return nil
}

type CloseGameParams struct {
	AcceptHead
	PlayerName string `json:"playerName" example:"test"` // 玩家在运营商系统中的唯一标识
}

func (o *CloseGameParams) CheckParams() error {
	if o.PlayerName == "" {
		return errors.New("playerName is empty")
	}
	return nil
}

func (o *OpenHandle) CloseGame() error {
	if err := o.Check(); err != nil {
		return err
	}
	o.Res = "ok"
	return nil
}

// HistoryListParams 记录列表请求
type HistoryListParams struct {
	AcceptHead
	PlayerName    string `json:"playerName" example:"test"`   // 玩家在运营商系统中的唯一标识
	GameID        int    `json:"gameId" example:"1"`          // 游戏id
	TransactionId string `json:"transactionId" example:"tx1"` // 供应商交易id
	Date          string `json:"date" example:"2020-01-01"`   // 日期
	Size          int    `json:"size" example:"10"`           // 每页条数 最大100
	Page          int    `json:"page" example:"1"`            // 页码
}

func (o *HistoryListParams) CheckParams() error {
	if o.Size > 100 {
		return errors.New("size up to 100")
	}
	return nil
}

// History 记录列表响应
type History struct {
	PlayerName    string  `json:"playerName" example:"test"`     // 玩家在运营商系统中的唯一标识
	GameId        int     `json:"gameId" example:"1"`            // 游戏id
	TransactionId string  `json:"transactionId" example:"tx1"`   // 供应商交易id
	AgentTxid     string  `json:"agentTxid" example:"tx2"`       // 运营商交易id
	Currency      string  `json:"currency" example:"USD"`        // 货币类型
	WinAmount     float64 `json:"winAmount" example:"1.00"`      // 赢得金额
	Bet           float64 `json:"bet" example:"10.00"`           // 下注金额
	PlayerAmount  float64 `json:"playerAmount" example:"800.85"` // 玩家余额(返奖后)
	Status        int     `json:"status" example:"1"`            // 状态 1:InProgress 2:Completed 3:CompleteInProcess 9:Canceled 10:CancelInProcess
}

func (o *OpenHandle) HistoryList(params *HistoryListParams) error {
	if err := o.Check(); err != nil {
		return err
	}
	var (
		historys []*History
		txns     []*business.Txn
	)
	q := global.GVA_DB.Model(&business.Txn{})
	if params.PlayerName != "" {
		q = q.Where("player_name = ?", params.PlayerName)
	}
	if params.GameID != 0 {
		q = q.Where("game_id = ?", params.GameID)
	}
	if params.TransactionId != "" {
		q = q.Where("txn_id = ?", params.TransactionId)
	}
	if params.Date != "" {
		t, err := time.Parse("2006-01-02", params.Date)
		if err == nil {
			q = q.Where("created_at >= ? and created_at < ?", t, t.AddDate(0, 0, 1))
		}
	}
	if params.Size < 0 {
		params.Size = 10
	} else if params.Size > 100 {
		params.Size = 100
	}
	if params.Page < 1 {
		params.Page = 1
	}
	q = q.Where("merchant_id = ?", o.Merchant.ID)
	err := q.Order("id desc").Limit(params.Size).Offset((params.Page - 1) * params.Size).Find(&txns).Error
	if err != nil {
		return err
	}
	for _, txn := range txns {
		historys = append(historys, &History{
			PlayerName:    txn.PlayerName,
			GameId:        txn.GameId,
			TransactionId: txn.TxnId,
			AgentTxid:     txn.PlatformTxnId,
			Currency:      txn.Currency,
			WinAmount:     helper.Div100(txn.Win),
			Bet:           helper.Div100(txn.Bet + txn.Raise),
			PlayerAmount:  helper.Div100(txn.AfterBal),
			Status:        int(txn.Status),
		})
	}
	o.Res = historys
	return nil
}

func (o *OpenHandle) WriteLog(err error) {
	var (
		remark = ""
		res    = response.Response{
			Code: 0,
			Data: o.Res,
			Msg:  "ok",
		}
		status = enum.Yes
	)
	if err != nil {
		remark = err.Error()
		res.Code = response.ERROR
		res.Msg = remark
		status = enum.No
	}
	if o.Merchant == nil {
		o.Merchant = &business.Merchant{}
	}
	s, _ := global.Json.Marshal(res)
	q, _ := global.Json.Marshal(o.Params)
	log := business.ApiLog{
		MerchantId: o.Merchant.ID,
		Type:       2,
		Agent:      o.Merchant.Agent,
		Url:        o.Url,
		Method:     "POST",
		Request:    string(q),
		Response:   string(s),
		Status:     uint8(status),
		Remark:     remark,
	}
	err = global.GVA_DB.Create(&log).Error
	if err != nil {
		global.GVA_LOG.Error("写api日志失败", zap.Error(err), zap.Any("log", log))
	}
	return
}
