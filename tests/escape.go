package main

import "log"

/*
# Go 的变量逃逸
所谓逃逸分析（Escape analysis）是指由编译器决定内存分配的位置，不需要程序员指定。

在函数中申请一个新的对象：
* 如果分配在栈中，则函数执行结束可自动将内存回收；
* 如果分配在堆中，则函数执行结束可交给 GC（垃圾回收）处理;

终端运行命令查看逃逸分析日志：`go build -gcflags=-m`

## 逃逸分析的意义
>* 逃逸分析的好处是为了减少gc的压力，不逃逸的对象分配在栈上，当函数返回时就回收了资源，不需要gc标记清除。
>* 逃逸分析完后可以确定哪些变量可以分配在栈上，栈的分配比堆快，性能好(逃逸的局部变量会在堆上分配 ,而没有发生逃逸的则有编译器在栈上分配)。
>* 同步消除，如果你定义的对象的方法上有同步锁，但在运行时，却只有一个线程在访问，此时逃逸分析后的机器码，会去掉同步锁运行。

## 一些问答
传递指针相比值传递减少了底层拷贝，可以提高效率，但是拷贝的数据量较小，由于指针传递会产生逃逸，可能会使用堆，也可能增加gc的负担，所以指针传递不一定是高效的。

## reference
* http://alblue.cn/articles/2020/07/07/1594131614114.html#toc_h2_11
*/

func main() {
	// 指针逃逸
	// StuRegist("三喜", 22)
	// 栈空间不足逃逸
	escapeStackNotEnough()
	// 动态类型逃逸。因 Println 方法的参数是 interface{} 类型
	log.Println("test end...")
	// 闭包引用逃逸
	f1 := escapeClosure()
	f1()
}

// 指针逃逸 start
type Stu struct {
	Name string
	Age  uint8
}

func StuRegist(name string, age uint8) *Stu {
	s1 := &Stu{
		Name: name,
		Age:  age,
	}
	return s1
}

// 指针逃逸 end

// 栈空间不足逃逸 start
func escapeStackNotEnough() {
	arr := make([]int, 10000, 10000)
	for index, _ := range arr {
		arr[index] = index + 1
	}
}

// 栈空间不足逃逸 end
// 闭包引用逃逸 start
func escapeClosure() func() int {
	v1, v2 := 0, 1
	return func() int {
		v1, v2 = v2, v1+v2
		return v1
	}
}

// 闭包引用逃逸 end
