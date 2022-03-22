package substringWithConcatenationOfAllWords

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	got := findSubstring("barfoothefoobarman", []string{"foo", "bar"})
	expected := []int{0, 9}
	if !reflect.DeepEqual(got, expected) {
		t.Error("err 01")
		return
	}
	t.Log("--end--")
}

func TestSlice1(t *testing.T) {
	s1 := "barfoothefoobarman"
	if s1[0:2] != "ba" {
		t.Error("err 001")
		return
	}
	if s1[1:2] != "ar" {
		t.Error("err 001")
		return
	}
	t.Log("--end--")
}

func TestDeepEqual1(t *testing.T) {
	if !IntSliceEqual([]int{0, 1}, []int{0, 1}) {
		t.Error("err 01")
		return
	}
	t.Log("--end--")
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for i, item := range a {
		if b[i] != item {
			return false
		}
	}

	return true
}
