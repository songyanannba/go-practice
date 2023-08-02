package backend

import (
	"github.com/lonng/nano"
	"github.com/lonng/nano/serialize/json"
	"github.com/lonng/nano/session"
	"github.com/pkg/errors"
	"slot-server/global"
	"slot-server/utils/helper"
)

var ClusterConnPool = map[string]*Connector{}

func RunCluster() error {
	listen := global.GVA_CONFIG.System.BackendAddr
	if listen == "" {
		return errors.Errorf("game listen address cannot empty")
	}

	masterAddr := global.GVA_CONFIG.System.MasterAddr
	if listen == "" {
		return errors.Errorf("master address cannot empty")
	}

	global.GVA_LOG.Println("Current backend server listen address", listen)
	global.GVA_LOG.Println("Remote master server address", masterAddr)

	session.Lifetime.OnClosed(OnSessionClosed)

	nano.Listen(listen,
		nano.WithAdvertiseAddr(masterAddr),
		nano.WithComponents(Component),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithDebugMode(),
		nano.WithLogger(global.GVA_LOG),
	)
	return nil
}

func CreateClusterConn(name, url, path, scheme string) *Connector {
	defer helper.PanicRecover()
	connector := NewConnector()
	ch := make(chan struct{})
	connector.OnConnected(func() {
		ch <- struct{}{}
	})
	connector.Start(url, path, scheme)
	<-ch
	ClusterConnPool[name] = connector
	//if helper.WaitTimeout(ch, 15) {
	//	return nil, errors.New("timeout")
	//}
	return connector
}
