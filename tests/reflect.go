package main

import (
	"fmt"
	"reflect"
)

/**
## 反射
### 1.有多个方法挂载一个结构体上
当有多个方法挂在结构体上，我们可以通过反射的方式，获取到对应的方法名，方法等信息

## reference
* 获取方法信息时，方法必须是导出的 https://segmentfault.com/q/1010000021676380

*/

type GreeterA struct {
	Name string
}

func (GreeterA) Api1() uint {
	fmt.Println("this is api1 of GreeterA")
	return 1
}

func (GreeterA) Api2() uint {
	fmt.Println("this is api2 of GreeterA")
	return 2
}

//通过反射的方式，获取挂在结构体下的方法信息
func getInfo(greeter interface{}) string {
	greeterValue := reflect.ValueOf(greeter)
	greeterType := reflect.TypeOf(greeter)
	//反射方式获取结构体下的方法的数量
	//注意：方法必须是导出的
	funcNum := greeterType.NumMethod()
	for i := 0; i < funcNum; i++ {
		tmpFuncName := greeterType.Method(i).Name
		tmpMethod := greeterValue.MethodByName(tmpFuncName)
		//获取处理的 handler
		caller := tmpMethod.Interface()
		tmpType := reflect.TypeOf(caller)
		// 方法调用时，需要组装好参数
		inputParamNum := tmpType.NumIn()
		inputParam := make([]reflect.Value, inputParamNum)
		fmt.Printf("num in: %s\n", tmpType.NumIn())
		// 结构体下的方法的调用
		res1 := reflect.ValueOf(caller).Call(inputParam)
		fmt.Printf("execute res: %v\n", res1)
	}
	// 这里的作用是？
	if greeterValue.Kind() == reflect.Ptr {
		greeterValue = greeterValue.Elem()
	}
	return greeterValue.FieldByName("Name").Interface().(string)
}

func main() {
	info := getInfo(&GreeterA{
		Name: "hello",
	})

	fmt.Printf("%v", info)
}
