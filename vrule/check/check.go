package check

import (
	"github.com/tidwall/gjson"
)

const (
	RuleTagRequired = "required" //
	RuleTagNotEmpty = "notEmpty" // notEmpty
	RuleTagStrNum   = "strNum"   //
	RuleTagInteger  = "integer"  //
	RuleTagFloat    = "float"    //
	RuleTagGt       = "gt"       // gt:1
	RuleTagGte      = "gte"      // gte:1
	RuleTagLt       = "lt"       // lt:1
	RuleTagLte      = "lte"      // lte:1
	RuleTagEq       = "eq"       // eq:1
	RuleTagNeq      = "neq"      // neq:1
	RuleTagIn       = "in"       // in:1,2,3
	RuleTagBetween  = "between"  // between:1,3
	RuleTagLength   = "length"   // length:1
	RuleTagMin      = "min"      // min:1
	RuleTagMax      = "max"      // max:10
	RuleTagEmail    = "email"    // email
	RuleTagPhone    = "phone"    // phone
	RuleTagDate     = "date"     // date
	RuleTagTime     = "time"     // time
	RuleTagUnique   = "unique"   // unique
)

// 内置message
var builtinMessages = map[string]string{
	RuleTagRequired: "{field} is required to not be empty",
	RuleTagNotEmpty: "{field} cannot be empty",
	RuleTagStrNum:   "{field} value must be an string number",
	RuleTagInteger:  "{field} value must be an integer",
	RuleTagFloat:    "{field} value must be a float number",
	RuleTagGt:       "{field} value must be greater than the field %s",
	RuleTagGte:      "{field} value should be greater or equal to the field %s",
	RuleTagLt:       "{field} value should be less than the field %s",
	RuleTagLte:      "{field} value should be less than or equal to the field %s",
	RuleTagIn:       "{field} value must be in the field %s",
	RuleTagEq:       "{field} value must be equal the field %s",
	RuleTagNeq:      "{field} value cannot be equal to the field %s",
	RuleTagBetween:  "{field} value must between %s",
	RuleTagLength:   "{field} value length must be %s",
	RuleTagMin:      "{field} value must be greater than or equal to %s",
	RuleTagMax:      "{field} value must be less than or equal to %s",
	RuleTagEmail:    "{field} value must be a valid email address",
	RuleTagPhone:    "{field} value must be a valid phone number",
	RuleTagDate:     "{field} value must be a valid date (YYYY-MM-DD)",
	RuleTagTime:     "{field} value must be a valid time (HH:MM:SS)",
	RuleTagUnique:   "{field} array elements must be unique",
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
