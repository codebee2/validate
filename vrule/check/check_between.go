package check

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

type CheckBetween struct{}

func (own *CheckBetween) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	tagValues := strings.Split(ruleTagInfo.TagValue, ",")
	v1, v2 := tagValues[0], tagValues[1]
	switch fieldVal.Type {
	case gjson.Number:
		v1f, _ := strconv.ParseFloat(v1, 64)
		v2f, _ := strconv.ParseFloat(v2, 64)
		return fieldVal.Float() >= v1f && fieldVal.Float() <= v2f
	case gjson.String:
		return fieldVal.String() >= v1 && fieldVal.String() <= v2
	}
	return false
}

func (own *CheckBetween) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := fmt.Sprintf(builtinMessages[ruleTagInfo.TagKey], ruleTagInfo.TagValue)
	return strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
}

func (own *CheckBetween) Tag() string {
	return RuleTagBetween
}
