package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"slot-server/global"
	"slot-server/model/business"
	"slot-server/pbs"
	"slot-server/utils"
	"slot-server/utils/helper"
	"strconv"
	"strings"
	"time"
)

func GetRecordMenuKey(userId uint, req *pbs.RecordMenuReq) string {
	return "{record_menu}:" + strconv.Itoa(int(userId)) + ":" + strconv.Itoa(int(req.GameId)) + ":" + req.Date
}

// GetRecordMenuCache  获取用户机台游戏记录缓存
func GetRecordMenuCache(userId uint, req *pbs.RecordMenuReq) (recordMenu *RecordMenu, err error) {
	key := GetRecordMenuKey(userId, req)
	recordMenu = &RecordMenu{}
	var result []byte
	result, err = global.GVA_REDIS.Get(context.Background(), key).Bytes()
	if err != nil {
		global.GVA_LOG.Error("GetRecordMenuCache err", zap.Error(err))
		return
	}
	err = json.Unmarshal(result, &recordMenu)
	if err != nil {
		global.GVA_LOG.Error("GetRecordMenuCache err", zap.Error(err))
		return
	}
	return
}

// SetRecordMenuCache  设置用户机台游戏记录缓存
func SetRecordMenuCache(userId uint, req *pbs.RecordMenuReq, recordMenu *RecordMenu) (err error) {
	key := GetRecordMenuKey(userId, req)
	var data []byte
	data, err = json.Marshal(recordMenu)
	err = global.GVA_REDIS.Set(context.Background(), key, data, 1*time.Hour).Err()
	return
}

type RecordMenu struct {
	DayHourMap map[int32][]int32 `json:"day_hour_map"`
	MaxTime    string            `json:"max_time"`
}

func GetRecordMenu(userId uint, req *pbs.RecordMenuReq) (dayHourMap map[int32][]int32, err error) {
	recordMenu := &RecordMenu{}
	recordMenu, err = GetRecordMenuCache(userId, req)
	if err != nil || recordMenu.MaxTime == "" {
		//无缓存查数据库
		recordMenu, err = DbGetRecordMenu(userId, req, "")
		if err != nil {
			return
		}
		//更新缓存
		err = SetRecordMenuCache(userId, req, recordMenu)
		if err != nil {
			global.GVA_LOG.Error("SetRecordMenuCache err", zap.Error(err))
		}
		dayHourMap = recordMenu.DayHourMap
		return
	}

	dayHourMap = recordMenu.DayHourMap
	var time1 = convertToServerTime(req).Format("2006-01")
	var time2 = time.Now().Format("2006-01")
	//如果不是当前月份,且有缓存返回缓存
	if time1 != time2 {
		return
	}

	//如果是当前月份,且最后一小时有数据,返回缓存
	nowTime := TimeOffset(time.Now(), req.TimeZone, false)
	if ContainsValue(dayHourMap[int32(nowTime.Day())], int32(nowTime.Hour())) {
		return
	}

	//如果是当前月份,且最后一小时没有数据,更新缓存
	addRecordMenu := &RecordMenu{}
	minTime := time.Time{}
	minTime, err = time.Parse("2006-01-02 15:04:05", recordMenu.MaxTime)
	if err != nil {
		global.GVA_LOG.Error("time.Parse err", zap.Error(err))
	}
	addRecordMenu, err = DbGetRecordMenu(userId, req, minTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		return
	}
	for k, v := range addRecordMenu.DayHourMap {
		for _, i2 := range v {
			recordMenu.DayHourMap[k] = append(recordMenu.DayHourMap[k], i2)
		}
	}
	recordMenu.MaxTime = addRecordMenu.MaxTime
	err = SetRecordMenuCache(userId, req, recordMenu)
	if err != nil {
		global.GVA_LOG.Error("SetRecordMenuCache err", zap.Error(err))
	}
	return
}

func ContainsValue(arr []int32, value int32) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func DbGetRecordMenu(userId uint, req *pbs.RecordMenuReq, minTime string) (recordMenu *RecordMenu, err error) {
	var (
		startTime time.Time
		endTime   time.Time
		layout    = "2006-01"
		times     []string
	)
	recordMenu = &RecordMenu{
		DayHourMap: make(map[int32][]int32),
	}
	if minTime == "" {
		startTime, err = time.ParseInLocation(layout, req.Date, time.FixedZone(req.TimeZone, 0))
		startTime = TimeOffset(startTime, req.TimeZone, false)
	} else {
		startTime, err = time.ParseInLocation("2006-01-02 15:04:05", minTime, time.FixedZone(req.TimeZone, 0))
		startTime = TimeOffset(startTime, req.TimeZone, true).Add(time.Second)
	}

	if err != nil {
		return
	}

	endTime = startTime.AddDate(0, 1, 0).Add(-time.Second)

	err = global.GVA_READ_DB.Model(&business.SlotRecord{}).
		Select("DATE_FORMAT(CONVERT_TZ(created_at, '+00:00', '"+GetOffsetString(req.TimeZone)+"'), '%d %H') as 'hour'").
		Where("user_id = ? ", userId).
		Where("slot_id = ? ", req.GameId).
		Where("created_at >= ? and created_at <= ?", startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05")).
		Group("hour").
		Scan(&times).Error
	global.GVA_LOG.Info("times", zap.Strings("times", times))
	maxDate := 0
	maxHour := 0
	for _, t := range times {
		day, hourStr, _ := strings.Cut(t, " ")
		date, _ := strconv.Atoi(day)
		hour, _ := strconv.Atoi(hourStr)
		recordMenu.DayHourMap[int32(date)] = append(recordMenu.DayHourMap[int32(date)], int32(hour))
		if date > maxDate {
			maxDate = date
			maxHour = hour
		} else if date == maxDate && hour > maxHour {
			maxHour = hour
		}
	}
	if maxDate == 0 {
		recordMenu.MaxTime = minTime
		return
	}
	recordMenu.MaxTime = time.Date(startTime.Year(), startTime.Month(), maxDate, maxHour+1, 0, 0, 0, startTime.Location()).Format("2006-01-02 15:04:05")

	return
}

func TimeOffset(date time.Time, timeZone string, flip bool) time.Time {
	offset := GetOffsetInt(timeZone)
	if flip {
		offset = -offset
	}
	return date.Add(time.Duration(offset) * time.Second)
}

func GetOffsetInt(timeZone string) int {
	// 加载目标时区//Asia/Shanghai
	global.GVA_LOG.Info("timeZone", zap.String("timeZone", timeZone))
	targetLocation, err := time.LoadLocation(timeZone)
	if err != nil {
		global.GVA_LOG.Error("无法加载目标时区:", zap.Error(err))
		return 0
	}

	now := time.Now()
	targetTime := now.In(targetLocation)

	// 计算目标时区与服务器本地时区的偏移量
	_, offset := targetTime.Zone()
	_, nowOffset := now.Zone()

	return offset - nowOffset
}

func GetOffsetString(timeZone string) string {
	offset := GetOffsetInt(timeZone)
	// 将时区偏移量格式化为"-07:00"字符串
	return fmt.Sprintf("%s%02d:%02d", getOffsetSign(offset), offset/3600, (offset%3600)/60)
}

// 根据偏移量的正负返回符号
func getOffsetSign(offset int) string {
	if offset < 0 {
		return "-"
	}
	return "+"
}

func convertToServerTime(req *pbs.RecordMenuReq) time.Time {
	layout := "2006-01"
	startTime, err := time.ParseInLocation(layout, req.Date, time.FixedZone(req.TimeZone, 0))
	if err != nil {
		global.GVA_LOG.Error("convertToServerTime err", zap.Error(err))
	}
	startTime = TimeOffset(startTime, req.TimeZone, false)
	return startTime
}

//---------------------------以下为列表缓存---------------------------------------------

func GetRecordListKey(userId uint, gameId int32, dateTime string) string {
	return "record:list:" + strconv.Itoa(int(userId)) + ":" + strconv.Itoa(int(gameId)) + ":" + dateTime
}

func GetRecordListCache(userId uint, gameId int32, dateTime string) (recordList []*pbs.RecordInfo, err error) {
	key := GetRecordListKey(userId, gameId, dateTime)
	recordList = []*pbs.RecordInfo{}
	err = utils.GetCache(key, &recordList)
	if err != nil {
		return nil, err
	}
	return
}

func SetRecordListCache(userId uint, gameId int32, dateTime string, recordList []*pbs.RecordInfo) (err error) {
	key := GetRecordListKey(userId, gameId, dateTime)
	randSec := helper.RandInt(120)
	randSec += 60
	err = utils.SetCache(key, recordList, time.Duration(randSec)*time.Second)
	if err != nil {
		return err
	}
	return
}

func DeleteRecordListCache(userId uint, gameId int32, dateTime string) (err error) {
	err = utils.DelCache(GetRecordListKey(userId, gameId, dateTime))
	if err != nil {
		return err
	}
	return
}

func DeleteRecordListCacheList(keys []string) (err error) {
	err = utils.DelCache(keys...)
	if err != nil {
		return err
	}
	return
}
