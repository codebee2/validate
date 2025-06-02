package vrule

import (
	"errors"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"sync"
)

type checkFunc func(ruleTagInfo *RuleTag, gRes gjson.Result) bool

var lk sync.Mutex
var BakedInValidators = map[string]checkFunc{
	"required": CkRequired, // 是否存在该字段
	"number":   CkNumber,   // number  // 校验 [1,"1"]
	"integer":  CkInteger,
}

type RuleTag struct {
	FieldKey string
	Tag      string // gt
	Value    string // gt->value
}

// gt:1
// required
// return {tag:gt,value:1}
func ParseRuleTag(info string) *RuleTag {
	result := &RuleTag{
		Tag:   info,
		Value: "",
	}
	firstIndex := strings.Index(info, ":")
	if firstIndex < 0 {
		return result
	}
	result.Tag = info[:firstIndex]
	result.Value = info[firstIndex+1:]
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
	_, err := strconv.ParseInt(gRes.Get(ruleTagInfo.Tag).String(), 10, 64)
	return err == nil
}

func CkInteger(ruleTagInfo *RuleTag, gRes gjson.Result) bool {
	val := gRes.Get(ruleTagInfo.FieldKey)
	// 如果有. 则表示是 float
	if strings.Contains(val.Raw, ".") {
		return false
	}
	return val.Type == gjson.Number
}

func CkEgt(ruleTagInfo *RuleTag, gRes gjson.Result) bool {
	val := gRes.Get(ruleTagInfo.FieldKey)
	// 如果有. 则表示是 float
	// 数字 999 123456
	// 字符串 999 123456
	// 字符串和 数字的比较性质完全不一样
	if strings.Contains(val.Raw, ".") {
		return false
	}
	return val.Type == gjson.Number
}
