package kubernetes

import (
	"gorm.io/gorm"
	"time"
)

type Cluster struct {
	ID         uint           `json:"id" gorm:"not null;unique;primary_key"`
	Name       string         `json:"name" form:"name" gorm:"comment:集群名称"`
	KubeConfig string         `gorm:"type:longText" json:"kube_config" form:"kube_config" gorm:"comment:kube_config"`
	ApiServer  string         `gorm:"type:longText" json:"api_server"  form:"api_server" gorm:"comment:master api address"`
	Prometheus string         `gorm:"type:longText" json:"prometheus" form:"prometheus" gorm:"comment:prometheus address"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Cluster) TableName() string {
	return "k8s_clusters"
}
