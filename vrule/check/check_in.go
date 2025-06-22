package check

import (
	"github.com/tidwall/gjson"
	"strings"
)

type CheckIn struct{}

func (own *CheckIn) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	switch fieldVal.Type {
	case gjson.Number:
		//ruleTagInfo.TagValue
	case gjson.String:
	}

	return false
}

func (own *CheckIn) CheckMsg(ruleTagInfo *RuleTag) string {
	return strings.Replace(builtinMessages[ruleTagInfo.TagKey], "{field}", ruleTagInfo.TagKey, 1)
}

func (own *CheckIn) Tag() string {
	return RuleTagIn
}
