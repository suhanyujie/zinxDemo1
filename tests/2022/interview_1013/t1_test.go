package interview_1013

import (
	"encoding/json"
	"fmt"
	"runtime"
	"testing"
)

type Student struct {
	Name string
}

func TestT1(t *testing.T) {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)

	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
	fmt.Println(str2)
	// [a b new]
	// [a b new]
	// [b new z x y]
}

func TestT2(t *testing.T) {
	fmt.Println(&Student{
		Name: "LiuDeHua",
	} == &Student{Name: "LiuDeHua"})

	fmt.Println(Student{
		Name: "LiuDeHua",
	} == Student{Name: "LiuDeHua"})
}

// Q3
func TestT3(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

// Q4
type StudentQ4 struct {
	Age int
}

func TestT4(t *testing.T) {
	kv := map[string]StudentQ4{"user1": StudentQ4{Age: 19}}
	// not compile
	// kv["user1"].Age = 20

	s := []StudentQ4{{Age: 19}}
	s[0].Age = 21
	fmt.Println(kv, s)
}

// Q5
type PeopleQ5 struct {
	name string `json:"name"`
}

func TestT5(t *testing.T) {
	json1 := `{
	"name": "user1"
}`
	var p1 PeopleQ5
	err := json.Unmarshal([]byte(json1), &p1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("people: ", p1)
	fmt.Printf("people: %v \n", p1)
	// people: {}
}

func TestT6(t *testing.T) {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
	// 可能正常打印，也可能 panic
}

func TestQ6(t *testing.T) {
	a, b := 1, 1
	defer func(x, y int) {
		fmt.Println(x + y)
	}(a, b)

	a, b = 2, 2
	defer func() {
		fmt.Println(a + b)
	}()

	a, b = 3, 3
	fmt.Println(a + b)
	// 6
	// 6
	// 2
}

func TestQ8(t *testing.T) {
	s1 := []int{0, 1, 2, 3}
	m1 := make(map[int]*int)
	for k, v := range s1 {
		m1[k] = &v
	}
	for k, v := range m1 {
		fmt.Println(k, "->", *v)
	}
	//0 -> 3
	//1 -> 3
	//2 -> 3
	//3 -> 3
}

// 考察 go 中的零值
// https://jingwei.link/2018/06/30/golang-interface-analysis.html
func TestQ9(t *testing.T) {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

func TestQ10(t *testing.T) {
	m1 := make(map[string]int)
	delete(m1, "h1")
	fmt.Println(m1["h1"])
}
