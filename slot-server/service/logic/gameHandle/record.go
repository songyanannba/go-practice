package gameHandle

import (
	"github.com/lonng/nano/session"
	"go.uber.org/zap"
	"slot-server/enum"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/service/cache"
	"time"
)

func RecordList(s *session.Session, userId uint, req *pbs.RecordListReq, ack *pbs.RecordListAck) (err error) {
	var (
		records         []business.SlotRecord
		recordList      []*pbs.RecordInfo
		startTime       time.Time
		endTime         time.Time
		layout          string = "2006-01-02 15:04:05"
		excludeRecordId        = s.Uint(enum.SessionDataGameRecordId)
	)
	startTime, err = time.ParseInLocation(layout, req.Time, time.FixedZone(req.TimeZone, 0))
	startTime = cache.TimeOffset(startTime, req.TimeZone, true)
	if err != nil {
		global.GVA_LOG.Error("time.ParseInLocation err", zap.Error(err))
		return
	}
	recordList, err = cache.GetRecordListCache(userId, req.GameId, startTime.Format("2006-01-02 15"))
	if err != nil || len(recordList) == 0 {
		endTime = startTime.Add(time.Hour)
		// 创建db
		err = global.GVA_READ_DB.Model(&business.SlotRecord{}).
			Select("created_at,id,currency,bet,gain,after_bal ").
			Where("user_id = ? ", userId).
			Where("slot_id = ? ", req.GameId).
			Where("created_at >= ? and created_at < ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).
			Order("`id` desc").
			Omit("ack").
			Find(&records).Error
		if err != nil {
			return err
		}

		for _, record := range records {
			if record.ID == excludeRecordId {
				continue
			}
			recordList = append(recordList, &pbs.RecordInfo{
				Time:    record.CreatedAt.Unix(),
				No:      business.FmtSlotRecordNo(record.ID),
				Uint:    record.Currency,
				Bet:     int64(record.Bet),
				Win:     int64(record.Gain),
				Balance: record.AfterBal,
			})
		}
	}
	ack.List = recordList
	err = cache.SetRecordListCache(userId, req.GameId, startTime.Format("2006-01-02 15"), recordList)
	if err != nil {
		global.GVA_LOG.Error("cache.SetRecordListCache err", zap.Error(err))
	}
	return nil
}

func RecordMenu(userId uint, req *pbs.RecordMenuReq, ack *pbs.RecordMenuAck) error {
	ack.List = make([]*pbs.RecordMenuDate, 0)
	listMap, err := cache.GetRecordMenu(userId, req)
	if err != nil {
		return err
	}

	for day, hour := range listMap {
		ack.List = append(ack.List, &pbs.RecordMenuDate{
			Date: day,
			Hour: hour,
		})
	}
	return nil
}
