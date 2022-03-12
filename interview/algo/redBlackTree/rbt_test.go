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
	tree.Insert("user4", "userName1")
	tree.Insert("user1", "userName1")
	tree.Insert("user2", "userName1")
	tree.Insert("user3", "userName1")
	tree.Insert("user5", "userName1")
	tree.Insert("user6", "userName1")
	tree.rootNode.Print(20)

	matchedNode := tree.Search("user5")
	matchedNode.Print(10)
}

func TestSearch1(t *testing.T) {
	tree := NewRBTree(Cmp1)
	tree.Insert("user4", "userName1")
	tree.Insert("user1", "userName1")
	tree.Insert("user2", "userName1")
	tree.Insert("user3", "userName1")
	tree.Insert("user5", "userName1")
	tree.Insert("user6", "userName1")

	matchedNode := tree.Search("user5")
	matchedNode.PrintOne()
}
