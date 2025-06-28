package check

import (
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type CheckLength struct{}

func (own *CheckLength) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if ruleTagInfo == nil {
		return false
	}
	// 只支持数组类型
	if !fieldVal.IsArray() {
		return false
	}
	length, err := strconv.Atoi(ruleTagInfo.TagValue)
	if err != nil {
		return false
	}
	return len(fieldVal.Array()) == length
}

func (own *CheckLength) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagLength]
	msg = strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
	return strings.Replace(msg, "%s", ruleTagInfo.TagValue, 1)
}

func (own *CheckLength) Tag() string {
	return RuleTagLength
}
