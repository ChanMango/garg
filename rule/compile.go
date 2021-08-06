package rule

import "regexp"

type RuleCompile string

var (
	AndMatch = "and|&"
	OrMatch  = "or|\\|"
	//err express
	Err_Require_Match        = "[^(required)]+" //required 只能单独存在
	Err_IlegelExpress_Mahtch = "(<!+)|(>!+)|(><)|(=<)|(=>)|(!<)|(!>)|(=!)|(<{2,})|(>{2,})|(={3,})+"
	Err_IlegelWord_Match     = "[^0-9a-zA-z+-><=!|&\\(\\)]+"
)
var ErrorRuleVerityList = []RuleCompile{RuleCompile(Err_Require_Match), RuleCompile(Err_IlegelExpress_Mahtch), RuleCompile(Err_IlegelWord_Match)}

func (c RuleCompile) Compile() (bool, Express) {

}
