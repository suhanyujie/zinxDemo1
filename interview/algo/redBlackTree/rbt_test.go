package redBlackTree

import "testing"

func TestNewRBTree(t *testing.T) {
	cmp := func(a, b interface{}) int {
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
	tree := NewRBTree(cmp)
}
