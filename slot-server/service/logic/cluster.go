package logic

import (
	"context"
	"errors"
	"google.golang.org/protobuf/proto"
	"slot-server/global"
	"slot-server/pbs"
	"slot-server/utils/env"
)

func PublishCluster(typ int, data string) error {
	if typ == 0 {
		return errors.New("cluster operate is nil")
	}
	operate := &pbs.ClusterOperate{
		Publisher: env.IP,
		Type:      int32(typ),
		Data:      data,
	}
	ctx := context.Background()
	msg, err := proto.Marshal(operate)
	if err != nil {
		return err
	}
	err = global.GVA_REDIS.Publish(ctx, "cluster", string(msg)).Err()
	return err
}
