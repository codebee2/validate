package validate

import (
	"fmt"
	"testing"
)

func Test_Rule(t *testing.T) {
	// 给你一个json数据
	data := `{"id":"1"}`

	vvv := NewValidator(data)
	vvv.AddRules(MS{
		"id": "required|integer",
	})
	vvv.AddMsg(MS{
		"id.required": "id是必须的",
		"id.integer":  "id必须是数字",
	})
	fmt.Println(vvv.Check())
	fmt.Println(vvv.GetErrorOne())
}
