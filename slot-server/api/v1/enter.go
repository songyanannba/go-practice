package v1

import (
	"slot-server/api/v1/business"
	"slot-server/api/v1/example"
	"slot-server/api/v1/system"
	"slot-server/api/v1/test"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	TestApiGroup     test.ApiGroup
	BusinessApiGroup business.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
