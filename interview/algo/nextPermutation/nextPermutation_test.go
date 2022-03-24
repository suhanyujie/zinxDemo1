package nextPermutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 7, 4, 3, 1}
	nextPermutation(nums)
	expected := []int{1, 3, 1, 2, 4, 7}
	if !reflect.DeepEqual(nums, expected) {
		t.Error("error 01")
		return
	}

	nums = []int{8, 5, 3, 2, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 5, 8}) {
		t.Error("error 02")
		return
	}
	t.Log("--end--")
}
