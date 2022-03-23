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

	t.Log("--end--")
}
