package business

import (
	"slot-server/global"
	"slot-server/utils/helper"
)

type Table interface {
	TableName() string
}

func Get[V any, T helper.Int](id T) (V, error) {
	var v V
	err := global.GVA_DB.First(&v, id).Error
	return v, err
}

func GetList[V any](where ...any) ([]V, error) {
	var v []V
	q := global.GVA_DB
	if len(where) > 0 {
		q = q.Where(where)
	}
	err := q.Find(&v).Error
	return v, err
}
