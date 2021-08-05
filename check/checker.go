package check

//
type Result map[string]string

func Check(val interface{}) (bool, Result) {
	return true, nil
}

type CheckerFunc = func(val interface{}) bool

func CustomChecker(value interface{}, ckFun CheckerFunc) bool {
	return ckFun(value)
}
