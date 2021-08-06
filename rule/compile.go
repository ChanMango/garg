package rule

import (
	"errors"
	"fmt"
	"git.xiaojukeji.com/chenyeung/garg/common"
	"regexp"
	"strings"
)

type RuleCompile string

var (
	And_OP_Match             RuleCompile = "\\band\\b|&"
	Or_OP_Match              RuleCompile = "\\bor\\b|\\|"
	In_OP_Match              RuleCompile = "\\bin\\b"
	NotIn_OP_Match           RuleCompile = "(\\bnot\\b|!)\\s*(in)"
	NormalExpress_Match      RuleCompile = "(>=)|(<=)|<|>|=|(!=)"
	CollectElemExpress_Match RuleCompile = "[a-z|A-Z|0-9]+" //专用于in ｜not in 获取元素
	//err express
	Err_Require_Match        RuleCompile = "[&|\\||>|<|=]+.*(required)+" //required 只能单独存在
	Err_IlegelExpress_Mahtch RuleCompile = "(<!+)|(>!+)|(><)|(=<)|(=>)|(!<)|(!>)|(=!)|(<{2,})|(>{2,})|(={3,})+"
	Err_IlegelWord_Match     RuleCompile = "[^0-9a-zA-z+-><=!|&\\(\\)\\s]+" //
	Err_HasAndOR_Match       RuleCompile = "(and|&)+.*(or|\\|)+|(or|\\|)+.*(and|&)+"
)
var ErrorRuleVerityList = []RuleCompile{Err_Require_Match, Err_IlegelExpress_Mahtch, Err_IlegelWord_Match, Err_HasAndOR_Match}

func ParseRule(val string) (Express, error) {
	andCmp, _ := regexp.Compile(string(And_OP_Match))
	orCmp, _ := regexp.Compile(string(Or_OP_Match))
	hasAnd := andCmp.MatchString(val)
	hasOr := orCmp.MatchString(val)
	bucket := NewCalBucket()
	if hasAnd {
		and_split_rule := andCmp.Split(val, -1)
		for _, sr := range and_split_rule {
			express, err := Parse(sr)
			if err != nil {
				fmt.Println("arg 解析规则时，出现内部错误", err.Error())
				return nil, err
			}
			bucket.andBucket = append(bucket.andBucket, express)
		}
		return bucket, nil
	}
	if hasOr {
		or_split_rule := orCmp.Split(val, -1)
		for _, sr := range or_split_rule {
			express, err := Parse(sr)
			if err != nil {
				fmt.Println("arg 解析规则时，出现内部错误", err.Error())
				return nil, err
			}
			bucket.orBucket = append(bucket.orBucket, express)
		}
		return bucket, nil
	}
	//集合表达式
	inCmp, _ := regexp.Compile(string(In_OP_Match))
	notInCmp, _ := regexp.Compile(string(NotIn_OP_Match))
	hasIn := inCmp.MatchString(val)
	hasNotIi := notInCmp.MatchString(val)
	collExpression := NewCollectionExpression()
	if hasIn {
		collExpression.op = IN_OperatorType
		eles := inCmp.FindStringSubmatch(val)
		collExpression.elems = append(collExpression.elems, eles[:])
		return collExpression, nil
	}
	if hasNotIi {
		collExpression.op = NI_OperatorType
		eles := notInCmp.FindStringSubmatch(val)
		collExpression.elems = append(collExpression.elems, eles[:])
		return collExpression, nil
	}
	//普通表达式
	ncp, _ := regexp.Compile(string(NormalExpress_Match))
	hasNormal := ncp.MatchString(val)
	if hasNormal {
		expression := NewNormalExpression()
		opStr := ncp.FindString(val)
		operateType := String2OperateType(opStr)
		expression.Op = operateType
		expression.expected = strings.TrimSpace(ncp.ReplaceAllString(val, ""))
		return expression, nil
	}
	//todo require的规则待实现

	return nil, errors.New("表达式为空，或表达式内容规则不合法")
}

func (c RuleCompile) Match(val string) (bool, error) {
	return regexp.MatchString(string(c), val)
}
func VerifyRuleIsLegal(rule string) (bool, error) {
	//先检查 表达式是否合法
	for _, rp := range ErrorRuleVerityList {
		match, _ := rp.Match(rule)
		if match {
			return false, common.Ilegle_Expression_Error
		}
	}
	return true, nil
}
