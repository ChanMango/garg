package common

import (
	"errors"
	"log"
	"runtime"
)

// err msg key tpl

const
(
	And_OPErrKey = "and_express_result"
	OR_OPErrKey  = "or_express_result"
)

var (
	NotStructTypeError      = errors.New("Not Support Type,only support struct type")
	Ilegle_Expression_Error = errors.New("表达式语法有误，请检查")
)

func PanicReportDecorator(f func() error) error {

	if err := recover(); err != nil {
		buf := make([]byte, 1<<16)
		stackSize := runtime.Stack(buf, true)
		log.Fatalf("_soda_fc_board_panic||errmsg=panic error: %s, %s", err, string(buf[0:stackSize]))
	}
	return f()
}
