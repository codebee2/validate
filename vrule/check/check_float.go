package check

import (
	"strings"

	"github.com/tidwall/gjson"
)

type CheckFloat struct{}

func (own *CheckFloat) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if fieldVal.Type != gjson.Number {
		return false
	}
	// 如果有. 则表示是 float
	return strings.Contains(fieldVal.Raw, ".")
}

func (own *CheckFloat) CheckMsg(ruleTagInfo *RuleTag) string {
	return strings.Replace(builtinMessages[ruleTagInfo.TagKey], "{field}", ruleTagInfo.TagKey, 1)
}

func (own *CheckFloat) Tag() string {
	return RuleTagFloat
}
