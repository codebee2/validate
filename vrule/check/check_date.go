package check

import (
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type CheckDate struct{}

func (own *CheckDate) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if fieldVal.Type != gjson.String {
		return false
	}
	_, err := time.Parse("2006-01-02", fieldVal.String())
	return err == nil
}

func (own *CheckDate) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagDate]
	return strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
}

func (own *CheckDate) Tag() string {
	return RuleTagDate
}
