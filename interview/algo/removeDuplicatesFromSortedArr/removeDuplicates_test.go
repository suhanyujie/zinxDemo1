package removeDuplicatesFromSortedArr

import "testing"

func TestRemoveDuplicates1(t *testing.T) {
	nums := []int{1, 1, 2, 4, 5, 6, 6, 7}
	res := removeDuplicates(nums)
	expected := 6
	if res != expected {
		t.Error("exec error 1")
		return
	}
	t.Log("--ok, end--")
}
