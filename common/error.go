package common

import "errors"

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
