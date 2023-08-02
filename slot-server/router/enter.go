package router

import (
	"slot-server/router/business"
	"slot-server/router/example"
	"slot-server/router/public"
	"slot-server/router/system"
	"slot-server/router/test"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Test     test.RouterGroup
	Business business.RouterGroup
	Public   public.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
