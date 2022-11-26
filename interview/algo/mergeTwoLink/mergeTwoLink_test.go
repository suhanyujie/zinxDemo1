package mergeTwoLink

import (
	"testing"
	"zinx_demo1/interview/other/jsonCode/jsonx"
)

func NewNode(v int) *ListNode {
	return &ListNode{
		Val:  v,
		Next: nil,
	}
}

func TestMerge1(t *testing.T) {
	list1 := []int{1, 3, 5, 8}
	list2 := []int{2, 4, 7, 9}

	expectList := []int{1, 2, 3, 4, 5, 7, 8, 9}
	expectLink := GetLink(expectList)
	link1 := GetLink(list1)
	link2 := GetLink(list2)
	t.Logf("link2: %v", jsonx.ToJsonIgnoreError(link2))
	res := Merge(link1, link2)
	t.Logf("res: %v", jsonx.ToJsonIgnoreError(res))
	t.Logf("expected: %v", jsonx.ToJsonIgnoreError(expectLink))
}

func GetLink(list []int) *ListNode {
	tmpWrap := &ListNode{}
	curNode := tmpWrap
	for _, v := range list {
		curNode.Next = NewNode(v)
		curNode = curNode.Next
	}
	return tmpWrap.Next
}
