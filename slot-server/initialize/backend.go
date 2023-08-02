package initialize

import (
	"slot-server/global"
	"slot-server/service/cluster/backend"
)

func RunBackend() {
	if global.GVA_DB != nil && global.GVA_CONFIG.System.Migrate {
		// 初始化表
		RegisterTables(global.GVA_DB)
	}
	if global.GVA_CONFIG.System.ConnectCluster {
		RunCluster()
	}
	return
}

func RunCluster() {
	for _, cluster := range global.GVA_CONFIG.System.Clusters {
		go backend.CreateClusterConn(
			cluster.Name,
			cluster.GetUrl(global.GVA_CONFIG.System.GateAddr),
			global.GVA_CONFIG.System.WsPath,
			cluster.WsScheme,
		)
	}
}
