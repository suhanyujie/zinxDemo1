package mergeSort

import "testing"

func TestMergeSort1(t *testing.T) {
	nums := []int{12, 98, 10, 8, 19, 27, 17, 20}
	result := mergeSort(nums)
	expected := []int{8, 10, 12, 17, 19, 20, 27, 98}
	for i, val := range expected {
		if result[i] != val {
			t.Errorf("sort error... val: %d", val)
		}
	}
	t.Log("--ok, end--")
}
