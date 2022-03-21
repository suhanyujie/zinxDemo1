package t1

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Stu1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJson1(t *testing.T) {
	json1 := "{\"name\": \"Samuel\", \"age\": 21}"
	stu1 := Stu1{}
	// 只有导出的 field 才能进行反序列化
	json.Unmarshal([]byte(json1), &stu1)
	fmt.Printf("%+v", stu1)
}
