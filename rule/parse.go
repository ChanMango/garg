package rule

//一期只实现了一个层级的计算 如 >10 and < 100 或者 >10 或者required
func Parse(ruleStr string) (Express, error) {
	legal, err := VerifyRuleIsLegal(ruleStr)
	if !legal {
		return nil, err
	}
	express, err := ParseRule(ruleStr)
	if err != nil {
		return nil, err
	}
	return express, nil
}
