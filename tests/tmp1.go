package main

import "fmt"

func main() {
	// map 内字段的匹配
	//res := IsLocalForum("东城")
	//fmt.Println(res)
	// 从 map 中删除一个元素
	// deleteEle()
	// 校验结构体是否实现了某个接口
	doTest1()
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
