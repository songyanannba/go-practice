package master

import (
	"fmt"
	"github.com/lonng/nano/benchmark/testdata"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/session"
)

var (
	Component = &component.Components{}
	Service   = newService()
)

func init() {
	Component.Register(Service)
}

type MasterService struct {
	component.Base
}

func newService() *MasterService {
	return &MasterService{}
}

func (ms *MasterService) Init() {}

func (ms *MasterService) AfterInit() {}

func (ms *MasterService) BeforeShutdown() {}

func (ms *MasterService) Shutdown() {}

func (ms *MasterService) Test(s *session.Session, data *testdata.Ping) error {
	fmt.Printf("data: %+v", *data)
	return s.Response(&testdata.Pong{Content: "master server pong"})
}
