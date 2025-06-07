package check

import (
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

type CheckStrNum struct{}

func (own *CheckStrNum) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	_, err := strconv.ParseInt(fieldVal.String(), 10, 64)
	return err == nil
}
func (own *CheckStrNum) CheckMsg(ruleTagInfo *RuleTag) string {
	return strings.Replace(builtinMessages[ruleTagInfo.TagKey], "{field}", ruleTagInfo.TagKey, 1)
}
func (own *CheckStrNum) Tag() string {
	return RuleTagStrNum
}
