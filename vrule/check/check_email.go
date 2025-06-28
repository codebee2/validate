package check

import (
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

type CheckEmail struct{}

var emailRegexp = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (own *CheckEmail) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if fieldVal.Type != gjson.String {
		return false
	}
	return emailRegexp.MatchString(fieldVal.String())
}

func (own *CheckEmail) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagEmail]
	return strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
}

func (own *CheckEmail) Tag() string {
	return RuleTagEmail
}
