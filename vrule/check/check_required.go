package check

import (
	"github.com/tidwall/gjson"
	"strings"
)

type CheckRequired struct{}

func (own *CheckRequired) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return fieldVal.Exists()
}
func (own *CheckRequired) CheckMsg(ruleTagInfo *RuleTag) string {
	return strings.Replace(builtinMessages[RuleTagRequired], "{field}", RuleTagRequired, 1)
}
func (own *CheckRequired) Tag() string {
	return RuleTagRequired
}
