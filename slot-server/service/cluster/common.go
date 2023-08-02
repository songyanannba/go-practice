package cluster

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/pbs"
	"slot-server/service/cluster/gate"
	"slot-server/utils/env"
	"slot-server/utils/helper"
	"time"
)

func SubscribeCluster() {
	ctx := context.Background()

	pubsub := global.GVA_REDIS.Subscribe(ctx, "cluster")
	defer func() {
		helper.PanicRecover()
		pubsub.Close()
	}()

	// 循环等待接收消息
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			global.GVA_LOG.Error("Error receiving redis receive message:", zap.Error(err))
			time.Sleep(time.Second * 5)
			continue
		}

		op := pbs.ClusterOperate{}
		err = proto.Unmarshal([]byte(msg.Payload), &op)
		if err != nil {
			global.GVA_LOG.Error("Error parse redis receive message:", zap.Error(err), zap.String("payload", msg.Payload))
			continue
		}
		OperateHandle(&op)
	}
}

func OperateHandle(op *pbs.ClusterOperate) {
	global.GVA_LOG.Info("cluster operate [publisher: " + op.Publisher +
		"], [type: " + helper.Itoa(op.Type) +
		"], [data: " + op.Data + "]")
	switch op.Type {
	case enum.ClusterOperateType1KickAccount:
		if env.IP == op.Publisher {
			return
		}
		uid := helper.Atoi(op.Data)
		if uid == 0 {
			return
		}
		gate.Service.Kick(int64(uid))
	}
}
