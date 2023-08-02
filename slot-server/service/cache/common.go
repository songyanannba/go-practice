package cache

import slotComponent "slot-server/service/slot/component"

func ClearLocalCache() {
	ClearSlotCache()
	slotComponent.FlushConfigCache()
}
