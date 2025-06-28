package check

import (
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type CheckTime struct{}

func (own *CheckTime) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if fieldVal.Type != gjson.String {
		return false
	}
	_, err := time.Parse("15:04:05", fieldVal.String())
	return err == nil
}

func (own *CheckTime) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagTime]
	return strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
}

func (own *CheckTime) Tag() string {
	return RuleTagTime
}
