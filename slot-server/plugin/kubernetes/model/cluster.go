package model

import (
	"gorm.io/gorm"
	"time"
)

type Cluster struct {
	ID                 uint           `json:"id" gorm:"not null;unique;primary_key"`
	Name               string         `json:"name" form:"name" gorm:"comment:集群名称"`
	KubeType           uint           `json:"kube_type" form:"kube_type" gorm:"comment:凭据类型1:KubeConfig,2:Token"`
	KubeConfig         string         `gorm:"type:longText" json:"kube_config" form:"kube_config" gorm:"comment:kube_config"`
	ApiAddress         string         `gorm:"type:longText" json:"api_address" form:"api_address" gorm:"comment:api_address"`
	PrometheusUrl      string         `gorm:"type:longText" json:"prometheus_url" form:"prometheus_url" gorm:"comment:prometheus 地址"`
	PrometheusAuthType uint           `json:"prometheus_auth_type" form:"prometheus_auth_type" gorm:"comment: 认证类型"`
	PrometheusUser     string         `gorm:"type:longText" json:"prometheus_user" form:"prometheus_user" gorm:"comment:用户名"`
	PrometheusPwd      string         `gorm:"type:longText" json:"prometheus_pwd" form:"prometheus_pwd" gorm:"comment:密码"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Cluster) TableName() string {
	return "k8s_clusters"
}
