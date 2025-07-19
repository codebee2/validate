package vrule

import (
	"errors"
	"strings"
	"sync"

	"github.com/codebee2/validate/vrule/check"
)

var lk sync.Mutex
var TagValidators = map[string]check.IRuleCheck{
	// required
	check.RuleTagRequired: &check.CheckRequired{}, // is existing field
	// type
	check.RuleTagStrNum:  &check.CheckStrNum{},  // str number
	check.RuleTagInteger: &check.CheckInteger{}, // integer
	// compare
	check.RuleTagBetween: &check.CheckBetween{},
	check.RuleTagGt:      &check.CheckGt{},
	check.RuleTagGte:     &check.CheckGte{},
	check.RuleTagLt:      &check.CheckLte{},
	check.RuleTagLte:     &check.CheckLte{},
	check.RuleTagEq:      &check.CheckEq{},
	check.RuleTagNeq:     &check.CheckNeq{},
	check.RuleTagIn:      &check.CheckIn{},
	check.RuleTagLength:  &check.CheckLength{},
	check.RuleTagMin:     &check.CheckMin{},
	check.RuleTagMax:     &check.CheckMax{},

	// other
	check.RuleTagDate:   &check.CheckDate{},
	check.RuleTagEmail:  &check.CheckEmail{},
	check.RuleTagPhone:  &check.CheckPhone{},
	check.RuleTagTime:   &check.CheckTime{},
	check.RuleTagUnique: &check.CheckUnique{},
}

// gt:1
// required
// return {tag:gt,value:1}
func ParseRuleTag(info string) *check.RuleTag {
	result := &check.RuleTag{
		TagKey:   info,
		TagValue: "",
	}
	firstIndex := strings.Index(info, ":")
	if firstIndex < 0 {
		return result
	}
	result.TagKey = info[:firstIndex]
	result.TagValue = info[firstIndex+1:]
	return result
}

func RegisterCheck(ck check.IRuleCheck) {
	lk.Lock()
	if _, ok := TagValidators[ck.Tag()]; ok {
		panic(errors.New("rule has exists"))
	}
	TagValidators[ck.Tag()] = ck
	lk.Unlock()
}
