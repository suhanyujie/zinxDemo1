package main

import "fmt"

func main() {
	test1()
}

// 位运算后，转为 byte 类型
func test1() {
	d1 := uint32(8)
	c1 := byte(d1 >> 4)
	c2 := c1 << 4
	fmt.Printf("%v\n", c2)
}
