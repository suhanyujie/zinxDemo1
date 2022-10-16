package t1

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/pkg/errors"
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

// 测试：带有时间的结构体，序列化成 json后，再转为 map[string]interface{} 是否还能识别成时间？
type Task1 struct {
	Title     string
	StartTime time.Time
}

func TestDe2Data(t *testing.T) {
	t1 := Task1{
		Title:     "t1",
		StartTime: time.Now(),
	}
	jsonStrByte, _ := json.Marshal(t1)
	map1 := make(map[string]interface{}, 3)
	json.Unmarshal(jsonStrByte, &map1)
	t.Log(map1)
}

func TestIfJudge1(t *testing.T) {
	var i interface{}
	if i == nil {
		t.Log("nil")
	} else {
		t.Log("not nil")
	}
}

func TestDefer1(t *testing.T) {
	res := defer1()
	fmt.Println(res)
}

func defer1() (res string) {
	defer func() {
		fmt.Println("step-1")
		res = "hello1"
	}()
	fmt.Println("step-2")
	res = "hello2"
	return
}

func TestCpuCoreNum1(t *testing.T) {
	fmt.Println(runtime.NumCPU())
}

func TestCtxForCancel(t *testing.T) {
	ctx := context.Background()
	curCtx, cancel := context.WithCancel(ctx)
	// eg: send message
	go sendMsg(curCtx)
	// eg: 查询用户
	time.Sleep(1 * time.Second)
	err := errors.New("test error")
	if err != nil {
		fmt.Println("[sendMsg] canceled")
		cancel()
		return
	}
	time.Sleep(3 * time.Second)
}

func TestCtxForDone(t *testing.T) {
	ctx := context.Background()
	curCtx, cancel := context.WithCancel(ctx)
	// eg: send message
	go sendMsg(curCtx)
	// eg: 查询用户
	time.Sleep(1 * time.Second)
	cancel()

	time.Sleep(3 * time.Second)
}

const AppTimeFormat = "2006-01-02 15:04:05"

func TestCtxForDeadline(t *testing.T) {
	ctx := context.Background()
	nowTime := time.Now()
	fmt.Println(nowTime.Format(AppTimeFormat))
	curCtx, cancel := context.WithDeadline(ctx, nowTime.Add(1*time.Second))
	// eg: send message
	go func() {
		if err := sendMsg(curCtx); err != nil {
			fmt.Println(err)
		}
	}()
	// eg: 查询用户
	time.Sleep(1 * time.Second)
	cancel()

	time.Sleep(3 * time.Second)
}

func sendMsg(ctx context.Context) error {
	t1, isOk := ctx.Deadline()
	if isOk {
		fmt.Println(t1.Format(AppTimeFormat))
	}
	time.Sleep(1 * time.Second)
	select {
	case dv := <-ctx.Done():
		fmt.Println("[sendMsg] done", dv)
		return ctx.Err()
	}
	// eg: do send msg
	time.Sleep(1 * time.Second)

	fmt.Println("[sendMsg] ok")

	return nil
}
