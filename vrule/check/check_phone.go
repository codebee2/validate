package check

import (
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

type CheckPhone struct{}

// 中国大陆手机号正则（以1开头，第二位3-9，后9位0-9，共11位）
var phoneRegexp = regexp.MustCompile(`^1[3-9]\d{9}$`)

func (own *CheckPhone) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	if fieldVal.Type != gjson.String {
		return false
	}
	return phoneRegexp.MatchString(fieldVal.String())
}

func (own *CheckPhone) CheckMsg(ruleTagInfo *RuleTag) string {
	msg := builtinMessages[RuleTagPhone]
	return strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
}

func (own *CheckPhone) Tag() string {
	return RuleTagPhone
}
