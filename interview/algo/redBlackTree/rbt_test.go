package redBlackTree

import "testing"

func Cmp1(a, b interface{}) int {
	aVal := a.(string)
	bVal := b.(string)
	if aVal < bVal {
		return -1
	} else if aVal > bVal {
		return 1
	} else {
		return 0
	}
}

func TestNewRBTree(t *testing.T) {
	tree := NewRBTree(Cmp1)
	if tree != nil {
		t.Error("tree is nil")
	}
	t.Log("--end--")
}

func TestInsert1(t *testing.T) {
	tree := NewRBTree(Cmp1)
	tree.Insert("user1", "userName1")

}
