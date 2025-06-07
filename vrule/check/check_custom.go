package check

import "github.com/tidwall/gjson"

type CheckCustom struct {
	TagName string
	CkFn    func(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool
	CkMsg   func(ruleTagInfo *RuleTag) string
}

func (own *CheckCustom) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return own.CkFn(ruleTagInfo, fieldVal)
}
func (own *CheckCustom) CheckMsg(ruleTagInfo *RuleTag) string {
	return own.CkMsg(ruleTagInfo)
}

func (own *CheckCustom) Tag() string {
	return own.TagName
}
