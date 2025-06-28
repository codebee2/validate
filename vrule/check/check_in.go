package check

import (
	"strings"

	"github.com/tidwall/gjson"
)

type CheckIn struct{}

func (own *CheckIn) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	values := strings.Split(ruleTagInfo.TagValue, ",")
	val := fieldVal.String()
	for _, v := range values {
		if val == v {
			return true
		}
	}
	return false
}

func (own *CheckIn) CheckMsg(ruleTagInfo *RuleTag) string {
	return strings.Replace(builtinMessages[ruleTagInfo.TagKey], "{field}", ruleTagInfo.TagKey, 1)
}

func (own *CheckIn) Tag() string {
	return RuleTagIn
}
