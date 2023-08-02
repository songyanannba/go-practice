package cache

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/samber/lo"
	"slot-server/model/business"
	"slot-server/pbs"
	"sort"
	"strconv"
)

var SlotCache *cmap.ConcurrentMap[string, *business.Slot]

func GetSlotKey(id uint) string {
	return strconv.Itoa(int(id))
}

func GetSlot(id uint) (res *business.Slot, err error) {
	if SlotCache == nil {
		if _, err = GetSlotList(); err != nil {
			return nil, err
		}
	}
	var ok bool
	res, ok = SlotCache.Get(GetSlotKey(id))
	if ok {
		return res, nil
	}

	return res, fmt.Errorf("id: %d not found", id)
}

func GetSlotList() (res []*business.Slot, err error) {
	if SlotCache != nil {
		return lo.MapToSlice(SlotCache.Items(), func(key string, value *business.Slot) *business.Slot {
			return value
		}), nil
	}

	res, err = business.GetList[*business.Slot]()
	if err != nil {
		return nil, err
	}

	m := lo.SliceToMap(res, func(value *business.Slot) (string, *business.Slot) {
		return GetSlotKey(value.ID), value
	})

	cache := cmap.New[*business.Slot]()
	cache.MSet(m)
	SlotCache = &cache
	return res, nil
}

func ClearSlotCache() {
	if SlotCache == nil {
		return
	}
	SlotCache = nil
}

func GetGameList() []*pbs.GameInfo {
	var (
		slots, _ = GetSlotList()
		gameList []*pbs.GameInfo
	)
	for _, slot := range slots {
		game := &pbs.GameInfo{
			Id:     int32(slot.ID),
			Name:   slot.NamePkg,
			Icon:   slot.Icon,
			Url:    slot.Url,
			Status: int32(slot.Status),
			Config: slot.ClientConf,
		}
		gameList = append(gameList, game)
	}
	sort.Slice(gameList, func(i, j int) bool {
		return gameList[i].Id < gameList[j].Id
	})
	return gameList
}
