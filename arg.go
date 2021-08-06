package garg

import (
	"git.xiaojukeji.com/chenyeung/garg/check"
)

type CheckIfc interface {
	Check(cf CheckerFunc) (bool, Result)
}

func Check(val interface{}) (bool, Result) {
	r := NewResult()
	isStruct, err := check.IsStruct(val)
	if !isStruct {
		r.Add("type", NewArgError(err))
		return false, r
	}
	//执行参数检查
	NewDefaultResolver(val).verify()
	return true, nil
}

type CheckerFunc = func(val interface{}) bool

func CustomChecker(value interface{}, ckFun CheckerFunc) bool {
	return ckFun(value)
}
