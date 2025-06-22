package validate

import (
	"fmt"
	"testing"
)

func Test_Rule(t *testing.T) {
	// 给你一个json数据
	data := `{"id":6,"data":{"b":1}}`

	//res := gjson.Get(data, "id")
	//fmt.Println(res.Type)
	//fmt.Println(res.Float())
	//return

	vvv := NewValidator(data)
	vvv.AddRules(MS{
		"id": "required|between:5,10",
	})
	vvv.AddMsg(MS{
		//"data.a.required": "data.a是必须的",
	})
	fmt.Println(vvv.Validate())
	fmt.Println(vvv.GetErrorOne())
}
