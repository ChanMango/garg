package rule

import (
	"errors"
	"fmt"
	"git.xiaojukeji.com/chenyeung/garg/check"
	"reflect"
	"regexp"
	"strings"
)

type RuleParser struct {
	ActualValue interface{}
	//ExpectValue interface{}
	Rule string
}

//只能通过该方法初始化
func NewParser(value reflect.Value, rule string) *RuleParser {
	return &RuleParser{
		ActualValue: value,
		Rule:        rule,
	}
}

//一期只实现了一个层级的计算 如 >10 and < 100 或者 >10 或者required
func (p *RuleParser) Parse() (Express, error) {
	//rule规则是否合理
	legal, err := VerifyRuleIsLegal(p.Rule)
	if !legal {
		return nil, err
	}
	//检查rule中的数据类型和actual是否一致
	express, err := internalParse(p.Rule, reflect.TypeOf(p.ActualValue))
	if err != nil {
		return nil, err
	}
	return express, nil
}

func internalParse(rule string, of reflect.Type) (Express, error) {
	andCmp, _ := regexp.Compile(string(And_OP_Match))
	orCmp, _ := regexp.Compile(string(Or_OP_Match))
	hasAnd := andCmp.MatchString(rule)
	hasOr := orCmp.MatchString(rule)
	bucket := NewCalBucket()
	if hasAnd {
		and_split_rule := andCmp.Split(rule, -1)
		for _, sr := range and_split_rule {
			express, err := internalParse(sr, of)
			if err != nil {
				fmt.Println("arg 解析规则时，出现内部错误", err.Error())
				return nil, err
			}
			bucket.andBucket = append(bucket.andBucket, express)
		}
		return bucket, nil
	}
	if hasOr {
		or_split_rule := orCmp.Split(rule, -1)
		for _, sr := range or_split_rule {
			express, err := internalParse(sr, of)
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
	eleCmp, _ := regexp.Compile(string(CollectElemExpress_Match))
	hasIn := inCmp.MatchString(rule)
	hasNotIi := notInCmp.MatchString(rule)
	collExpression := NewCollectionExpression()
	if hasIn {
		collExpression.op = IN_OperatorType
		eles := eleCmp.FindStringSubmatch(rule)
		container, err := check.CreateContainer(eles, of)
		if err != nil {
			return nil, err
		}
		collExpression.elems = container
		return collExpression, nil
	}
	if hasNotIi {
		collExpression.op = NI_OperatorType
		eles := eleCmp.FindStringSubmatch(rule)
		container, err := check.CreateContainer(eles, of)
		if err != nil {
			return nil, err
		}
		collExpression.elems = container
		return collExpression, nil
	}
	//普通表达式
	ncp, _ := regexp.Compile(string(NormalExpress_Match))
	hasNormal := ncp.MatchString(rule)
	if hasNormal {
		expression := NewNormalExpression()
		opStr := ncp.FindString(rule)
		operateType := String2OperateType(opStr)
		expression.Op = operateType
		expression.expected = strings.TrimSpace(ncp.ReplaceAllString(rule, ""))
		return expression, nil
	}
	rcp, _ := regexp.Compile(string(Required_OP_Match))
	hasNeed := rcp.MatchString(rule)
	if hasNeed {
		expression := NewNormalExpression()
		opStr := ncp.FindString(rule)
		operateType := String2OperateType(opStr)
		expression.Op = operateType
		//required 不需要 期望值，只需要判断实际值是否是默认值
		return expression, nil
	}

	return nil, errors.New("表达式为空，或表达式内容规则不合法-> " + rule)
}
