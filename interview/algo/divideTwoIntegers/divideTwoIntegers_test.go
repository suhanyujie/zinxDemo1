package divideTwoIntegers

import (
	"fmt"
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	if divide(10, 3) != 3 {
		t.Error("err 1")
		return
	}
	if divide(7, -3) != -2 {
		t.Error("err 2")
		return
	}
	if divide(-2147483648, -1) != 2147483647 {
		t.Error("err 3")
		return
	}

	t.Log("--end--")
	return
}

func TestOverflow1(t *testing.T) {
	var i1 int
	i1 = 2 << 31
	fmt.Printf("%d", i1)
	fmt.Printf("is greater: %v", i1 > math.MaxInt32)
}
