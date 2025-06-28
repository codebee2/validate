package check

import (
	"strings"

	"github.com/tidwall/gjson"
)

type CheckUnique struct{}

func (own *CheckUnique) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if !fieldVal.IsArray() {
		return false
	}
	arr := fieldVal.Array()
	seen := make(map[string]struct{}, len(arr))
	for _, v := range arr {
		s := v.String()
		if _, ok := seen[s]; ok {
			return false
		}
		seen[s] = struct{}{}
	}
	return true
}

func (own *CheckUnique) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagUnique]
	return strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
}

func (own *CheckUnique) Tag() string {
	return RuleTagUnique
}
