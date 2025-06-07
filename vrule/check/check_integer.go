package check

import (
	"github.com/tidwall/gjson"
	"strings"
)

type CheckInteger struct{}

func (own *CheckInteger) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if fieldVal.Type != gjson.Number {
		return false
	}
	// 如果有. 则表示是 float
	return !strings.Contains(fieldVal.Raw, ".")
}

func (own *CheckInteger) CheckMsg(ruleTagInfo *RuleTag) string {
	return strings.Replace(builtinMessages[ruleTagInfo.TagKey], "{field}", ruleTagInfo.TagKey, 1)
}

func (own *CheckInteger) Tag() string {
	return RuleTagInteger
}
