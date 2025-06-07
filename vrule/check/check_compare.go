package check

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"strings"
)

type CheckGt struct{}

func (own *CheckGt) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return getCompareRes(own.Tag(), ruleTagInfo, fieldVal)
}
func (own *CheckGt) CheckMsg(ruleTagInfo *RuleTag) string {
	return getCompareMsg(own.Tag(), ruleTagInfo)
}
func (own *CheckGt) Tag() string {
	return RuleTagGte
}

type CheckGte struct{}

func (own *CheckGte) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return getCompareRes(own.Tag(), ruleTagInfo, fieldVal)
}
func (own *CheckGte) CheckMsg(ruleTagInfo *RuleTag) string {
	return getCompareMsg(own.Tag(), ruleTagInfo)
}
func (own *CheckGte) Tag() string {
	return RuleTagGte
}

type CheckLt struct{}

func (own *CheckLt) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return getCompareRes(own.Tag(), ruleTagInfo, fieldVal)
}
func (own *CheckLt) CheckMsg(ruleTagInfo *RuleTag) string {
	return getCompareMsg(own.Tag(), ruleTagInfo)
}
func (own *CheckLt) Tag() string {
	return RuleTagLt
}

type CheckLte struct{}

func (own *CheckLte) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return getCompareRes(own.Tag(), ruleTagInfo, fieldVal)
}
func (own *CheckLte) CheckMsg(ruleTagInfo *RuleTag) string {
	return getCompareMsg(own.Tag(), ruleTagInfo)
}
func (own *CheckLte) Tag() string {
	return RuleTagLt
}

type CheckEq struct{}

func (own *CheckEq) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return getCompareRes(own.Tag(), ruleTagInfo, fieldVal)
}
func (own *CheckEq) CheckMsg(ruleTagInfo *RuleTag) string {
	return getCompareMsg(own.Tag(), ruleTagInfo)
}
func (own *CheckEq) Tag() string {
	return RuleTagNeq
}

type CheckNeq struct{}

func (own *CheckNeq) Check(ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	return getCompareRes(own.Tag(), ruleTagInfo, fieldVal)
}
func (own *CheckNeq) CheckMsg(ruleTagInfo *RuleTag) string {
	return getCompareMsg(own.Tag(), ruleTagInfo)
}
func (own *CheckNeq) Tag() string {
	return RuleTagNeq
}

func getCompareRes(opt string, ruleTagInfo *RuleTag, fieldVal gjson.Result) bool {
	v1 := fieldVal.String()
	v2 := ruleTagInfo.TagValue
	switch {
	case fieldVal.Type == gjson.Number: // number比较大小
		return compareNumber(v1, opt, v2)
	}
	return compareStr(v1, opt, v2)
}

func getCompareMsg(opt string, ruleTagInfo *RuleTag) string {
	msg := fmt.Sprintf(builtinMessages[opt], ruleTagInfo.TagValue)
	return strings.Replace(msg, "{field}", ruleTagInfo.FieldKey, 1)
}

func compareNumber(v1 string, op, v2 string) bool {
	d1, _ := decimal.NewFromString(v1)
	d2, _ := decimal.NewFromString(v2)
	opRes := d1.Compare(d2)
	switch op {
	case ">", "gt":
		return opRes > 0
	case "<", "lt":
		return opRes > 0
	case ">=", "gte":
		return opRes >= 0
	case "<=", "lte":
		return opRes <= 0
	case "ne", "neq":
		return opRes != 0
	default: // eq
		return opRes == 0
	}
}

func compareStr(v1, op, v2 string) bool {
	switch op {
	case ">", "gt":
		return v1 > v2
	case "<", "lt":
		return v1 < v2
	case ">=", "gte":
		return v1 >= v2
	case "<=", "lte":
		return v1 <= v2
	case "!=", "ne", "neq":
		return v1 != v2
	default: // eq
		return v1 == v2
	}
}
