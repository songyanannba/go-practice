package enum

import "errors"

const (
	Yes = 1 // 是
	No  = 2 // 否
)

const (
	CommonStatusUnknown    = 0  // 通用状态未知
	CommonStatusBegin      = 1  // 通用状态开始
	CommonStatusProcessing = 2  // 通用状态处理中
	CommonStatusClose      = 8  // 通用状态关闭
	CommonStatusError      = 9  // 通用状态异常
	CommonStatusFinish     = 10 // 通用状态完成
)

var (
	ErrUnknown        = errors.New("unknown")
	ErrBusy           = errors.New("Service_Busy")
	ErrNoMoney        = errors.New("No_Money")
	ErrNoNet          = errors.New("No_NET")
	ErrNoServer       = errors.New("No_Server")
	ErrSysError       = errors.New("Sys_Error")
	ErrTokenInvalid   = errors.New("Token_Invalid")
	ErrRecordNotExist = errors.New("Record_Not_Exist")
)
