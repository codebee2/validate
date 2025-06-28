package check

import (
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type CheckMax struct{}

func (own *CheckMax) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if ruleTagInfo == nil {
		return false
	}
	max, err := strconv.ParseFloat(ruleTagInfo.TagValue, 64)
	if err != nil {
		return false
	}

	switch fieldVal.Type {
	case gjson.Number:
		return fieldVal.Float() <= max
	case gjson.String:
		return float64(len(fieldVal.String())) <= max
	case gjson.JSON:
		if fieldVal.IsArray() {
			return float64(len(fieldVal.Array())) <= max
		}
	}
	return false
}

func (own *CheckMax) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagMax]
	msg = strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
	return strings.Replace(msg, "%s", ruleTagInfo.TagValue, 1)
}

func (own *CheckMax) Tag() string {
	return RuleTagMax
}
