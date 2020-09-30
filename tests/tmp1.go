package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// map 内字段的匹配
	//res := IsLocalForum("东城")
	//fmt.Println(res)
	// 从 map 中删除一个元素
	// deleteEle()
	// 校验结构体是否实现了某个接口
	// doTest1()
	//efunc()
	// test 引用传参
	// refFunc1()
	// 函数 len
	// testFuncLen()
	// 测试 switch
	//testSwitch()
	// 切片和切片指针
	sliceAndSlicePoint()
}

func IsLocalForum(needleName string) map[string]interface{} {
	initReturn := map[string]interface{}{
		"is_local":   0,
		"local_name": "",
	}
	localForumArr := map[string][]string{
		"北京市": []string{"东城"},
	}
	for _, arr := range localForumArr {
		for _, tmpName := range arr {
			if tmpName == needleName {
				initReturn["is_local"] = 1
				initReturn["local_name"] = tmpName
				return initReturn
			}
		}
	}

	return initReturn
}

// 内建类型数组元素的删除
func deleteEle() {
	a1 := map[string]string{
		"0": "su",
		"1": "han",
		"2": "yu",
	}
	delete(a1, "0")
	fmt.Println(a1)
}

// # 校验结构体是否实现了某个接口 start
// * 也称"接口守卫" interface guard
// ## reference https://www.reddit.com/r/golang/comments/j0b73d/what_is_var_type_type_used_for/
type Human interface {
	Say()
}

type Stu1 struct {
	Name string
}

type Stu2 struct {
	Name string
}

// todo
func (_this *Stu1) Say() {
	fmt.Printf("I'm a student 001: %s\n", _this.Name)
}

func doTest1() {
	var _ Human = &Stu1{}
	// var _ Human = &Stu2{}
	fmt.Println("ok")
	// 我们可以将 nil 转换为结构体指针
	s1 := (*Stu1)(nil)
	fmt.Printf("%v\n", s1)
}

// 校验结构体是否实现了某个接口 end

// 返回的异常类型是指针，实际返回是 nil start
type eP struct {
}

func (*eP) Error() string {
	return "err"
}

var (
	Err1 = errors.New("test error")
)

func efunc1() *eP {
	return nil
}

func efunc2() error {
	return efunc1()
	//var err error = efunc1()
	//err := efunc1()
	//log.Printf("%v, %v\n", err, err == nil)
	//return nil
	// err1 := err.(*eP)
	// log.Printf("%v,%v\n", err1, err1 == nil)
}

func efunc() error {
	v1 := efunc2()
	log.Printf("%v\n", v1)
	return nil
}

// 返回的异常类型是指针，实际返回是 nil end

// 引用传参 start
func refFunc1() {
	s1 := "苏汉"
	b1 := []byte(s1)
	log.Printf("%v\n", b1)
	refFunc2(&b1)
	log.Printf("%v\n", b1)
}

func refFunc2(b1 *[]byte) {
	*b1 = append(*b1, 10)
	*b1 = []byte("fafds")
}

// 函数 len
func testFuncLen() {
	s1 := "test1123123"
	log.Printf("string len is: %d\n", len(s1))
}

//test switch
func testSwitch() {
	t1 := 4
	switch t1 {
	case 1:
		log.Printf("1 val is: %v\n", t1)
	case 2:
		log.Printf("2 val is: %v\n", t1)
	case 3, 4:
		log.Printf("3 val is: %v\n", t1)
	}
}

// 切片和切片指针
func sliceAndSlicePoint() {
	s1 := []string{"hello", "world"}
	ps1 := []*[]string{}
	ps1 = append(ps1, &s1)

	log.Printf("%v\n", ps1)
}
