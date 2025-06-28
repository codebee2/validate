package check

import (
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type CheckMin struct{}

func (own *CheckMin) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if ruleTagInfo == nil {
		return false
	}
	min, err := strconv.ParseFloat(ruleTagInfo.TagValue, 64)
	if err != nil {
		return false
	}

	switch fieldVal.Type {
	case gjson.Number:
		return fieldVal.Float() >= min
	case gjson.String:
		return float64(len(fieldVal.String())) >= min
	case gjson.JSON:
		if fieldVal.IsArray() {
			return float64(len(fieldVal.Array())) >= min
		}
	}
	return false
}

func (own *CheckMin) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagMin]
	msg = strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
	return strings.Replace(msg, "%s", ruleTagInfo.TagValue, 1)
}

func (own *CheckMin) Tag() string {
	return RuleTagMin
}
