package validate

import (
	"fmt"
	"testing"
)

func Test_Rule(t *testing.T) {
	// 给你一个json数据
	data := `{"id":10}`

	vvv := NewValidator(data)
	vvv.AddRules(MS{
		"id": "required|integer|gte:10",
	})
	vvv.AddMsg(MS{
		"id.required": "id是必须的",
		"id.integer":  "id必须是数字",
		"id.gte":      "id必须大于10",
	})
	fmt.Println(vvv.Check())
	fmt.Println(vvv.GetErrorOne())
}
