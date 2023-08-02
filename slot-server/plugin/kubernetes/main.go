package kubernetes

import (
	"github.com/gin-gonic/gin"
	"slot-server/global"
	sysModel "slot-server/model/system"
	"slot-server/plugin/kubernetes/model"
	"slot-server/plugin/kubernetes/router"
	"slot-server/plugin/plugin-tool/utils"
)

type KubernetesPlugin struct {
}

func CreateKubernetesPlug() *KubernetesPlugin {
	global.GVA_DB.AutoMigrate(model.Cluster{})
	utils.RegisterApis(
		sysModel.SysApi{ApiGroup: "集群管理", Method: "POST", Path: "/kubernetes/clusters/list", Description: "列表"},
		sysModel.SysApi{ApiGroup: "集群管理", Method: "POST", Path: "/kubernetes/clusters/create", Description: "创建"},
		sysModel.SysApi{ApiGroup: "集群管理", Method: "POST", Path: "/kubernetes/clusters/getById", Description: "获取"},
		sysModel.SysApi{ApiGroup: "集群管理", Method: "PUT", Path: "/kubernetes/clusters/update", Description: "更新"},
		sysModel.SysApi{ApiGroup: "集群管理", Method: "DELETE", Path: "/kubernetes/clusters/delete", Description: "删除"},
		sysModel.SysApi{ApiGroup: "集群管理", Method: "DELETE", Path: "/kubernetes/clusters/deleteByIds", Description: "批量删除"},
		sysModel.SysApi{ApiGroup: "资源列表管理", Method: "POST", Path: "/kubernetes/proxy/*/*", Description: "创建"},
		sysModel.SysApi{ApiGroup: "资源列表管理", Method: "GET", Path: "/kubernetes/proxy/*/*", Description: "获取"},
		sysModel.SysApi{ApiGroup: "资源列表管理", Method: "PUT", Path: "/kubernetes/proxy/*/*", Description: "更新"},
		sysModel.SysApi{ApiGroup: "资源列表管理", Method: "DELETE", Path: "/kubernetes/proxy/*/*", Description: "删除"},
		sysModel.SysApi{ApiGroup: "资源列表管理", Method: "PATCH", Path: "/kubernetes/proxy/*/*", Description: "局部更新"},
		sysModel.SysApi{ApiGroup: "WebSocket", Method: "GET", Path: "/kubernetes/pods/terminal", Description: "终端"},
		sysModel.SysApi{ApiGroup: "WebSocket", Method: "GET", Path: "/kubernetes/pods/logs", Description: "终端日志"},
		sysModel.SysApi{ApiGroup: "普罗米修斯监控", Method: "POST", Path: "/kubernetes/metrics/get", Description: "监控数据"},
	)
	//utils.RegisterMenus(
	//	sysModel.SysBaseMenu{
	//		Name:      "Kubernetes",
	//		Path:      "kubernetes",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/index.vue",
	//		Sort:      1000,
	//		Meta:      sysModel.Meta{Title: "Kubernetes管理", Icon: "cloudy"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "clusters",
	//		Path:      "clusters",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/clusters/index.vue",
	//		Sort:      0,
	//		Meta:      sysModel.Meta{Title: "集群管理", Icon: "cloudy"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "nodes",
	//		Path:      "nodes",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/nodes/index.vue",
	//		Sort:      1,
	//		Meta:      sysModel.Meta{Title: "节点管理", Icon: "film"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "nodelabelstaints",
	//		Path:      "nodelabelstaints",
	//		Hidden:    true,
	//		Component: "plugin/kubernetes/view/nodelabelstaints/index.vue",
	//		Sort:      2,
	//		Meta:      sysModel.Meta{Title: "标签与污点", Icon: "star"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "workloads",
	//		Path:      "workloads",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/workloads/index.vue",
	//		Sort:      3,
	//		Meta:      sysModel.Meta{Title: "工作负载", Icon: "position"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "openkruise",
	//		Path:      "openkruise",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/openkruise/index.vue",
	//		Sort:      4,
	//		Meta:      sysModel.Meta{Title: "OpenKruise", Icon: "position"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "namespaces",
	//		Path:      "namespaces",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/namespaces/index.vue",
	//		Sort:      5,
	//		Meta:      sysModel.Meta{Title: "命名空间", Icon: "crop"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "networks",
	//		Path:      "networks",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/networks/index.vue",
	//		Sort:      6,
	//		Meta:      sysModel.Meta{Title: "网络管理", Icon: "ship"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "configs",
	//		Path:      "configs",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/configs/index.vue",
	//		Sort:      7,
	//		Meta:      sysModel.Meta{Title: "配置管理", Icon: "scale-to-original"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "storages",
	//		Path:      "storages",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/storages/index.vue",
	//		Sort:      8,
	//		Meta:      sysModel.Meta{Title: "存储管理", Icon: "coin"},
	//	},
	//	sysModel.SysBaseMenu{
	//		Name:      "accessControls",
	//		Path:      "accessControls",
	//		Hidden:    false,
	//		Component: "plugin/kubernetes/view/accessControls/index.vue",
	//		Sort:      9,
	//		Meta:      sysModel.Meta{Title: "访问控制", Icon: "key"},
	//	},
	//)
	return &KubernetesPlugin{}
}

func (*KubernetesPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitClusterRouter(group)
	router.RouterGroupApp.InitProxyRouter(group)
	router.RouterGroupApp.InitMetricRouter(group)
}

func (*KubernetesPlugin) RouterPath() string {
	return "kubernetes"
}
