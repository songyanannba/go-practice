package logic

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/protobuf/proto"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/system/request"
	"slot-server/pbs"
	"slot-server/service/cache"
	"slot-server/service/cluster/backend"
	"slot-server/service/system"
	"slot-server/utils"
	"slot-server/utils/helper"
)

func ParseBackendToken(head *pbs.ReqHead) (*request.BaseClaims, error) {
	if head == nil {
		return nil, errors.New("token invalidation")
	}
	if head.Token == enum.AdminDefaultToken {
		return &request.BaseClaims{
			ID:       0,
			Username: "system",
			NickName: "system",
		}, nil
	}
	s := system.JwtService{}
	if s.IsBlacklist(head.Token) {
		return nil, errors.New("您的帐户异地登陆或令牌失效")
	}
	j := utils.NewJWT()
	claims, err := j.ParseToken(head.Token)
	if err != nil {
		global.GVA_LOG.Error("从token中获取后台jwt解析信息失败: " + err.Error())
		return nil, err
	}
	return &claims.BaseClaims, err
}

func ReqBackendOperate(conn *backend.Connector, token string, typ int, data interface{}) (*pbs.BackendOperateAck, error) {
	if conn == nil {
		return nil, errors.New("cluster connection is nil")
	}
	// 请求游戏服务器清除缓存
	jsonData, _ := global.Json.Marshal(data)
	req := &pbs.BackendOperate{
		Head: &pbs.ReqHead{Token: token},
		Type: int32(typ),
		Data: string(jsonData),
	}
	res := &pbs.BackendOperateAck{}
	err := Req(conn, "BindService.BackendOperate", req, res)
	helper.RecordError(err, typ, data)
	return res, err
}

func Req(c *backend.Connector, route string, v proto.Message, ack proto.Message) (err error) {
	ch := make(chan struct{}, 1)
	err = c.Request(route, v, func(data interface{}) {
		err = proto.Unmarshal(data.([]byte), ack)
		ch <- struct{}{}
	})
	if err != nil {
		return
	}
	if helper.WaitTimeout(ch, 15) {
		return errors.New("timeout")
	}
	global.GVA_LOG.Skip(1).Println("ack :\n", spew.Sdump(ack))
	return
}

func Notify(c *backend.Connector, route string, v proto.Message) (err error) {
	err = c.Notify(route, v)
	return
}

func HandleBackendOperate(req *pbs.BackendOperate) error {
	switch req.Type {
	case enum.BackendOperateType1RefreshCache:
		cache.ClearLocalCache()
	}
	return nil
}
