package rule

import (
	"reflect"
	"regexp"
)

type OperatorType int
type RelationType int

var (
	//not !
	NOT_OperatorType OperatorType = -7
	//!= NE <>
	NE_OperatorType OperatorType = 3 //("not"+"="-> -7+10=0)
	//=
	EQ_OperatorType OperatorType = 10
	//<
	LT_OperatorType OperatorType = 1
	//>
	GT_OperatorType OperatorType = 2
	//<=
	LE_OperatorType OperatorType = 11
	//>=
	GE_OperatorType OperatorType = 12
	//need
	NEED_OperatorType OperatorType = 13
	//in
	IN_OperatorType OperatorType = 14
	NI_OperatorType OperatorType = 7
)
var (
	//运算关系提起
	And_RelationType RelationType = 0
	Or_RelationType  RelationType = 1
)

func (op OperatorType) String() string {
	msg := ""
	switch op {
	case LE_OperatorType:
		msg = "<"
	case GE_OperatorType:
		msg = ">"
	case NE_OperatorType:
		msg = "!="
	case EQ_OperatorType:
		msg = "="
	case LT_OperatorType:
		msg = "<="
	case GT_OperatorType:
		msg = ">="
	case NI_OperatorType:
		msg = "not in"
	case IN_OperatorType:
		msg = "in"
	}
	return msg
}

type RuleWrapper struct {
	op OperatorType
}

func (w RuleWrapper) Cal(actual interface{}) {
	switch w.op {
	case GE_OperatorType:

	}
}

//一期只实现了一个层级的计算 如 >10 and < 100 或者 >10 或者required
func Parse(ruleStr string) *RuleWrapper {
	and_mh, _ := regexp.MatchString("and|&", ruleStr)
	or_mh, _ := regexp.Compile("or|\\|")

	return &RuleWrapper{}
}
