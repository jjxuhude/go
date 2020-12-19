package loadgen

import (
	"bytes"
	"demo/loadgen/lib"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ParamSet struct {
	Caller     lib.Caller
	Lps        uint32
	DurationNS time.Duration
	TimeoutNS  time.Duration
	ResultCh   chan *lib.CallResult
}

func (pset *ParamSet) Check() error {
	var errMsgs []string
	if pset.Caller == nil {
		errMsgs = append(errMsgs, "无效的调用")
	}
	if pset.TimeoutNS == 0 {
		errMsgs = append(errMsgs, "无效的超时时间")
	}
	if pset.Lps == 0 {
		errMsgs = append(errMsgs, "无效的载荷量")
	}
	if pset.DurationNS == 0 {
		errMsgs = append(errMsgs, "无效的负载持续时间")
	}

	var buf bytes.Buffer
	buf.WriteString("检测参数......")
	if errMsgs != nil {
		errMsg := strings.Join(errMsgs, "")
		buf.WriteString(fmt.Sprintf("不能解析(%s)", errMsg))
		logger.Infoln(buf.String())
		return errors.New(errMsg)
	}
	buf.WriteString(fmt.Sprintf("通过：timeoutNs=%s lps=%d durationNs=%s",
		pset.TimeoutNS, pset.Lps, pset.DurationNS))
	logger.Infoln(buf.String())
	return nil
}
