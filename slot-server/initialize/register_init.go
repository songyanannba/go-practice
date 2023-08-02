package initialize

import (
	_ "slot-server/source/example"
	_ "slot-server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
