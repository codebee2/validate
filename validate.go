package validate

import (
	"encoding/json"
	"fmt"
	"github.com/codebee2/validate/vrule"
	"github.com/codebee2/validate/vrule/check"
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
func (own ErrorsField) Add(field, validator, message string) {
	if _, ok := own[field]; ok {
		own[field][validator] = message
	} else {
		own[field] = MS{validator: message}
	}
}

type Validator struct {
	sourceData      []byte
	validatorRule   MS           // 规则
	validatorMsg    MS           //错误消息
	parseRes        gjson.Result // sourceData解析结果
	RuleList        []check.IRuleCheck
	validatorErrors ErrorsField // 错误信息
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

func (own *Validator) Validate() bool {
	gjRes := gjson.ParseBytes(own.sourceData)
	var fieldVal gjson.Result
	for field, rules := range own.validatorRule {
		fieldVal = gjRes.Get(field)
		for _, rule := range strings.Split(rules, "|") {
			ruleTag := vrule.ParseRuleTag(rule)
			ruleTag.FieldKey = field
			ruleCheck, ok := vrule.TagValidators[ruleTag.TagKey]
			if !ok {
				own.validatorErrors.Add(field, ruleTag.TagKey, "invalid validator")
				return false
			}
			if !ruleCheck.Check(ruleTag, fieldVal) {
				msgKey := fmt.Sprintf(field + "." + ruleTag.TagKey)
				fmt.Println(msgKey, "======")
				errorMsg := ""
				if _, msgOk := own.validatorMsg[msgKey]; msgOk {
					errorMsg = own.validatorMsg[msgKey]
				} else {
					errorMsg = ruleCheck.CheckMsg(ruleTag)
				}
				own.validatorErrors.Add(field, msgKey, errorMsg)
				return false
			}
		}
	}
	return true
}
