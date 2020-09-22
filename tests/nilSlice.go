package main

import "fmt"

func main() {
	testNilSlice()
}

func testNilSlice() {
	var s1 []string
	var i1 int
	fmt.Printf("%v，isOk: %T\n", i1, i1)
	fmt.Printf("%v，isOk: %v\n", s1, s1 == nil)
}
