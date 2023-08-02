package service

import (
	"slot-server/service/business"
	"slot-server/service/example"
	"slot-server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
