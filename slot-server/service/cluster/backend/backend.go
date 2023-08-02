package backend

import (
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/examples/cluster/protocol"
	"github.com/lonng/nano/session"
	"github.com/pingcap/errors"
	"log"
)

var (
	Component = &component.Components{}
	Service   = newService()
)

func init() {
	Component.Register(Service)
}

func OnSessionClosed(s *session.Session) {
	Service.userDisconnected(s)
}

type BackendService struct {
	component.Base
	group *nano.Group
}

func newService() *BackendService {
	return &BackendService{
		group: nano.NewGroup("all-users"),
	}
}

func (bs *BackendService) Init() {}

func (bs *BackendService) AfterInit() {}

func (bs *BackendService) BeforeShutdown() {}

func (bs *BackendService) Shutdown() {}

func (bs *BackendService) userDisconnected(s *session.Session) {
	if err := bs.group.Leave(s); err != nil {
		log.Println("Remove user from group failed", s.UID(), err)
		return
	}
	log.Println("User session disconnected", s.UID())
}

type SyncMessage struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (bs *BackendService) SyncMessage(s *session.Session, msg *SyncMessage) error {
	if err := s.RPC("MasterService.Stats", &protocol.MasterStats{Uid: s.UID()}); err != nil {
		return errors.Trace(err)
	}
	return bs.group.Broadcast("onMessage", msg)
}

func s() {
}
