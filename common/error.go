package common

import "errors"

// err msg key tpl
var (
	And_OPErrKey = "and_express_result"
	OR_OPErrKey  = "or_express_result"
)
var NotStructTypeError = errors.New("Not Support Type,only support struct type")
