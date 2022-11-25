package reverseLink

import (
	"fmt"
	"testing"
)

func TestReverseLink1(t *testing.T) {
	link1 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	result := ReverseList(link1)
	fmt.Printf("result: %v", result)
}
