package validate

import (
	"fmt"
	"testing"
)

func Test_Rule(t *testing.T) {
	// 给你一个json数据
	data := `{"id":10,"data":{"b":1}}`

	vvv := NewValidator(data)
	vvv.AddRules(MS{
		"id":     "required|integer|gte:10",
		"data.a": "required|integer|gte:10",
	})
	vvv.AddMsg(MS{
		"data.a.required": "data.a是必须的",
	})
	fmt.Println(vvv.Validate())
	fmt.Println(vvv.GetErrorOne())
}
