package remvoeElement

import "testing"

func TestRemoveElement(t *testing.T) {
	if removeElement([]int{1, 1, 3}, 1) != 1 {
		t.Error("error 1")
		return
	}

	if removeElement([]int{1, 3, 3}, 3) != 1 {
		t.Error("error 1")
		return
	}
}
