package rule

import (
	"git.xiaojukeji.com/chenyeung/garg/common"
	"regexp"
)

type RuleCompile string

var (
	And_OP_Match             RuleCompile = "\\band\\b|&"
	Or_OP_Match              RuleCompile = "\\bor\\b|\\|"
	In_OP_Match              RuleCompile = "\\bin\\b"
	NotIn_OP_Match           RuleCompile = "(\\bnot\\b|!)\\s*(in)"
	NormalExpress_Match      RuleCompile = "(>=)|(<=)|<|>|=|(!=)"
	CollectElemExpress_Match RuleCompile = "(?:(?![(\\bin\\b)|(\\bnot\\s*\\bin\\b)]))\\w+" //专用于in ｜not in 获取元素
	Required_OP_Match        RuleCompile = "^(required)+(?![^(required)])|^(not\\s*null)+(?![^(not\\s%null)])"
	//err express
	Err_Require_Match        RuleCompile = "[&|\\||>|<|=]+.*(required)+" //required 只能单独存在
	Err_IlegelExpress_Mahtch RuleCompile = "(<!+)|(>!+)|(><)|(=<)|(=>)|(!<)|(!>)|(=!)|(<{2,})|(>{2,})|(={3,})+"
	Err_IlegelWord_Match     RuleCompile = "[^0-9a-zA-z+-><=!|&\\(\\)\\s]+" //
	Err_HasAndOR_Match       RuleCompile = "(and|&)+.*(or|\\|)+|(or|\\|)+.*(and|&)+"
)
var ErrorRuleVerityList = []RuleCompile{Err_Require_Match, Err_IlegelExpress_Mahtch, Err_IlegelWord_Match, Err_HasAndOR_Match}

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
