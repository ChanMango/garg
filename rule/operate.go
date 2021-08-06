package rule

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
	And_RelationType OperatorType = 0
	Or_RelationType  RelationType = 1
)

func OperateType2String(op OperatorType) string {
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
func String2OperateType(op string) (tp OperatorType) {
	switch op {
	case "<=":
		tp = LE_OperatorType
	case ">=":
		tp = GE_OperatorType
	case "!=":
		tp = NEED_OperatorType
	case "=":
		tp = EQ_OperatorType
	case "<":
		tp = LT_OperatorType
	case ">":
		tp = GT_OperatorType
	case "not in":
		tp = NI_OperatorType
	case "in":
		tp = IN_OperatorType
	}
	return
}
