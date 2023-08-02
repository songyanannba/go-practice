package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"slot-server/global"
	"slot-server/model/common/request"
	"slot-server/plugin/kubernetes/model"
)

type ClusterService struct{}

//@function: GetClustersInfoList
//@description: 分页获取列表
//@param: info request.PageInfo
//@return: list interface{}, total int64, err error

func (cls *ClusterService) GetClustersInfoList(cluster model.Cluster, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(model.Cluster{})
	var clsList []model.Cluster

	if cluster.Name != "" {
		db = db.Where("name LIKE ?", "%"+cluster.Name+"%")
	}

	err = db.Count(&total).Error

	if err != nil {
		return clsList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["name"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't matched any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return clsList, total, err
			}

			err = db.Order(OrderStr).Find(&clsList).Error
		} else {
			err = db.Order("id").Find(&clsList).Error
		}
	}

	return clsList, total, err
}

//@function: GetGlusterById
//@description: 根据id获取GetGlusterById
//@param: id float64
//@return: cluster kubernetes.Cluster, err error

func (cls *ClusterService) GetGlusterById(id int) (cluster model.Cluster, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cluster).Error
	return
}

//@function: CreateCluster
//@description: 新增基础cluster
//@param: cluster kubernetes.Cluster
//@return: err error

func (cls *ClusterService) CreateCluster(cluster model.Cluster) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ?", cluster.Name).First(&model.Cluster{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同集群")
	}
	return global.GVA_DB.Create(&cluster).Error
}

//@author: Eagle
//@function: UpdateCluster
//@description: 根据id更新Cluster
//@param: cluster kubernetes.Cluster
//@return: err error

func (cls *ClusterService) UpdateCluster(cluster model.Cluster) (err error) {
	return global.GVA_DB.Where("id = ?", cluster.ID).First(&model.Cluster{}).Updates(&cluster).Error
}

//@author: Eagle
//@function: DeleteCluster
//@description: 删除cluster
//@param: cluster kubernetes.Cluster
//@return: err error

func (cls *ClusterService) DeleteCluster(req request.GetById) (err error) {
	var cl model.Cluster
	if err = global.GVA_DB.Where("id = ?", req.ID).First(&cl).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err = global.GVA_DB.Delete(&cl).Error; err != nil {
		return err
	}

	return err
}

//@author: Eagle
//@function: DeleteClustersByIds
//@description: 删除选中集群
//@param: ids request.IdsReq
//@return: err error

func (cls *ClusterService) DeleteClustersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Cluster{}, "id in ?", ids.Ids).Error
	return err
}
