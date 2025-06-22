package check

import (
	"github.com/tidwall/gjson"
)

const (
	RuleTagRequired = "required" //
	RuleTagStrNum   = "strNum"   //
	RuleTagInteger  = "integer"  //
	RuleTagGt       = "gt"       // gt:1
	RuleTagGte      = "gte"      // gte:1
	RuleTagLt       = "lt"       // lt:1
	RuleTagLte      = "lte"      // lte:1
	RuleTagEq       = "eq"       // eq:1
	RuleTagNeq      = "neq"      // neq:1
	RuleTagIn       = "in"       // in:1,2,3
	RuleTagBetween  = "between"  // between:1,3
	RuleTagLength   = "length"   // length:1
)

// 内置message
var builtinMessages = map[string]string{
	RuleTagRequired: "{field} is required to not be empty",
	RuleTagStrNum:   "{field} value must be an string number",
	RuleTagInteger:  "{field} value must be an integer",
	RuleTagGt:       "{field} value must be greater than the field %s",
	RuleTagGte:      "{field} value should be greater or equal to the field %s",
	RuleTagLt:       "{field} value should be less than the field %s",
	RuleTagLte:      "{field} value should be less than or equal to the field %s",
	RuleTagIn:       "",
	RuleTagEq:       "{field} value must be equal the field %s",
	RuleTagNeq:      "{field} value cannot be equal to the field %s",
	RuleTagBetween:  "{field} value must between %s",
	RuleTagLength:   "",
}

type RuleTag struct {
	FieldKey string
	TagKey   string // gt
	TagValue string // gt->value
}

type IRuleCheck interface {
	Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool
	CheckMsg(ruleTagInfo *RuleTag) string
	Tag() string
}
