package t1

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Stu1 struct {
	Name  string            `json:"name"`
	Age   int               `json:"age"`
	Tools map[string]string `json:"tools"`
}

func TestJson1(t *testing.T) {
	json1 := "{\"name\": \"Samuel\", \"age\": 21}"
	stu1 := Stu1{}
	// 只有导出的 field 才能进行反序列化
	json.Unmarshal([]byte(json1), &stu1)
	fmt.Printf("%+v", stu1)
}

// 结构体能否判等。可以，但如果结构体里有复合类型，则不行
func TestEqual1(t *testing.T) {
	stu1 := Stu1{
		Name: "1",
	}
	stu2 := Stu1{
		Name: "1",
	}
	if reflect.DeepEqual(stu1, stu2) {
		t.Log("euqal")
		return
	}
	t.Log("--end--")
}

// 命名返回值，是否会自动初始化：并没有！
func TestNamedReturn1(t *testing.T) {
	res := NamedReturn1()
	t.Log(res)
}

func NamedReturn1() (m1 map[string]interface{}) {
	m1["name"] = 1
	return
}

type Stt1 struct {
	TableId int64 `json:"tableId"`
}

func TestUnmarshal1(t *testing.T) {
	m1 := map[string]interface{}{
		"tableId": "123",
	}
	data1 := Stt1{}
	jsonBytes, _ := json.Marshal(m1)
	json.Unmarshal(jsonBytes, &data1)
	t.Log(data1.TableId)
}
