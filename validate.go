package validate

import (
	"encoding/json"
	"fmt"
	"github.com/codebee2/validate/vrule"
	"github.com/tidwall/gjson"
	"strings"
)

type MS map[string]string

//	{
//		"field": {
//			"required": "error message",
//			"min_len": "error message1"
//		}
//	}
type ErrorsField map[string]MS

// Add 添加error
func (ef ErrorsField) Add(field, validator, message string) {
	if _, ok := ef[field]; ok {
		ef[field][validator] = message
	} else {
		ef[field] = MS{validator: message}
	}
}

type Validator struct {
	sourceData      []byte
	validatorRule   MS           // 规则
	validatorMsg    MS           //错误消息
	parseRes        gjson.Result // sourceData解析结果
	validatorErrors ErrorsField  // 错误信息
}

func (own *Validator) GetErrorOne() string {
	if len(own.validatorErrors) == 0 {
		return ""
	}
	for _, ms := range own.validatorErrors {
		for _, msg := range ms {
			return msg
		}
	}
	return ""
}

func (own *Validator) GeAllError() ErrorsField {
	d, _ := json.Marshal(own.validatorErrors)
	result := make(ErrorsField)
	_ = json.Unmarshal(d, &result)
	return result
}

// AddRules 添加规则
func (own *Validator) AddRules(rules MS) {
	own.validatorRule = rules
}

// AddMsg 添加错误消息
func (own *Validator) AddMsg(msg MS) {
	own.validatorMsg = msg
}

func NewValidator(data string) *Validator {
	return &Validator{
		sourceData:      []byte(data),
		parseRes:        gjson.Result{},
		validatorErrors: make(ErrorsField),
	}
}

func (own *Validator) Check() bool {
	gjRes := gjson.ParseBytes(own.sourceData)
	for field, rules := range own.validatorRule {
		for _, rule := range strings.Split(rules, "|") {
			// required|gte:5
			// // required|in:1,2,3,4
			// required|between:1,2,3
			ruleTag := vrule.ParseRuleTag(rule)
			ruleTag.FieldKey = field
			validatorsFn, ok := vrule.BakedInValidators[ruleTag.TagKey]
			if !ok {
				own.validatorErrors.Add(field, ruleTag.TagKey, "无效校验器")
				return false
			}
			if !validatorsFn(ruleTag, gjRes) {
				msgKey := fmt.Sprintf(field + "." + ruleTag.TagKey)
				own.validatorErrors.Add(field, msgKey, own.validatorMsg[msgKey])
				return false
			}
		}
	}
	return true
}
