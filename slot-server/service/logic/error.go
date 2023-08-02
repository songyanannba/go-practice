package logic

import (
	"fmt"
	"github.com/lonng/nano/session"
	"go.uber.org/zap"
	"path"
	"runtime"
	"slot-server/global"
	. "slot-server/pbs"
	"slot-server/utils"
	"time"
)

type Error struct {
	Code   Code
	Path   string
	Reason string
}

func NewErr(code Code, args ...interface{}) *Error {
	idaErr := &Error{}
	idaErr.Code = code
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	_, line := f.FileLine(pc)

	idaErr.Path = fmt.Sprintf("%s[%v:%v]", time.Now().Format("2006-01-02 15:04:05"), path.Base(f.Name()), line)
	for _, arg := range args {
		if val, ok := arg.(error); ok {
			idaErr.Reason += val.Error()
		} else {
			b, _ := global.Json.Marshal(arg)
			idaErr.Reason += string(b)
		}
	}

	return idaErr
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%v, reason=%v", e.Code, e.Reason)
}

// RespError 回应错误信息
func RespError(s *session.Session, req ReqHeads, ack AckHeads, err error) error {
	var (
		reqHead *ReqHead
		ackHead *AckHead
	)
	if req != nil {
		reqHead = req.GetHead()
	}
	if ack != nil {
		ackHead = ack.GetHead()
	}
	if reqHead == nil {
		reqHead = &ReqHead{}
	}
	if ackHead == nil {
		ackHead = &AckHead{}
	}
	ackHead.Uid = reqHead.Uid
	reqHead.Token = ""

	if ec, ok := err.(*Error); ok {
		ackHead.Code = ec.Code
		ackHead.Message = ec.Reason
	} else {
		ackHead.Message = err.Error()
	}
	global.GVA_LOG.WithOptions(zap.AddCallerSkip(1)).Error("\nrespError",
		zap.Int32("uid", reqHead.Uid),
		zap.String("playerName", reqHead.Platform),
		zap.String("ip", utils.RemoteIpString(s)),
		zap.String("req", fmt.Sprintf("%+v", req)),
		zap.String("ack", fmt.Sprintf("%+v", ack)),
		zap.Error(err),
	)

	return s.Response(ack)
}
