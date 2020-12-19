package lib

import "time"

type RawReq struct {
	ID  int64
	Req []byte
}

type RawResp struct {
	ID     int64
	Resp   []byte
	Err    error
	Elapse time.Duration
}

type RetCode int

const (
	RET_CODE_SUCCESS              RetCode = 0
	RET_CODE_WARNING_CALL_TIMEOUT         = 1001
	RET_COCE_ERROR_CALL                   = 2001
	RET_CODE_ERROR_RESPONSE               = 2002
	RET_CODE_ERROR_CALEE                  = 2003
	RET_CODE_FATAL_CALL                   = 3001
)

func GetRetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RET_CODE_SUCCESS:
		codePlain = "成功"
	case RET_CODE_WARNING_CALL_TIMEOUT:
		codePlain = "调用超时警告"
	case RET_COCE_ERROR_CALL:
		codePlain = "调用错误"
	case RET_CODE_ERROR_RESPONSE:
		codePlain = "响应内容错误"
	case RET_CODE_ERROR_CALEE:
		codePlain = "被测软件内部错误"
	case RET_CODE_FATAL_CALL:
		codePlain = "致命错误"
	default:
		codePlain = "未知错误"
	}
	return codePlain
}

type CallResult struct {
	ID     int64         //ID
	Req    RawReq        //原生请求
	Resp   RawResp       //原生响音
	Code   RetCode       //响应代码
	Msg    string        //结果成因描述
	Elapse time.Duration //耗时
}

const (
	STATUS_ORININAL uint32 = 0
	STATUS_STARTING uint32 = 1
	STATUS_STARTED  uint32 = 2
	STATUS_STOPPING uint32 = 3
	STATUS_STOPPED  uint   = 4
)

type Generator interface {
	Start() bool
	Stop() bool
	Status() uint32
	CallCount() int64
}
