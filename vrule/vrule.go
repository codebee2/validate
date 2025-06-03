package vrule

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"sync"
)

type checkFunc func(ruleTagInfo *RuleTag, gRes gjson.Result) bool

var lk sync.Mutex
var BakedInValidators = map[string]checkFunc{
	"required": CkRequired, // 是否存在该字段
	"number":   CkNumber,   // number  // 字符串1解析
	"integer":  CkInteger,
	"gte":      CkGte,
}

type RuleTag struct {
	FieldKey string
	TagKey   string // gt
	TagValue string // gt->value
}

// gt:1
// required
// return {tag:gt,value:1}
func ParseRuleTag(info string) *RuleTag {
	result := &RuleTag{
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

func RegisterRule(tag string, fn checkFunc) {
	lk.Lock()
	if _, ok := BakedInValidators[tag]; ok {
		panic(errors.New("rule has exists"))
	}
	BakedInValidators[tag] = fn
	lk.Unlock()
}

func CkRequired(ruleTagInfo *RuleTag, gRes gjson.Result) bool {
	return gRes.Get(ruleTagInfo.FieldKey).Exists()
}

// key存在则校验
func CkNumber(ruleTagInfo *RuleTag, gRes gjson.Result) bool {
	_, err := strconv.ParseInt(gRes.Get(ruleTagInfo.TagKey).String(), 10, 64)
	return err == nil
}

func CkInteger(ruleTagInfo *RuleTag, gRes gjson.Result) bool {
	val := gRes.Get(ruleTagInfo.FieldKey)
	if val.Type != gjson.Number {
		return false
	}
	// 如果有. 则表示是 float
	return !strings.Contains(val.Raw, ".")
}

func CkGte(ruleTagInfo *RuleTag, gRes gjson.Result) bool {
	val := gRes.Get(ruleTagInfo.FieldKey)
	v1 := val.String()
	v2 := ruleTagInfo.TagValue
	switch {
	case val.Type == gjson.Number: // number比较大小
		return compareNumber(v1, ">=", v2)
	}
	return compareStr(v1, ">=", v2)
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
