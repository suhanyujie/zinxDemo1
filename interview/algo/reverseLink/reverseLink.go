package reverseLink

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	curNode := pHead
	arr := make([]int, 0)
	for curNode != nil {
		arr = append(arr, curNode.Val)
		curNode = curNode.Next
	}
	if len(arr) < 1 {
		return pHead
	}
	newHead := &ListNode{}
	curNode = newHead
	for i := len(arr) - 1; i >= 0; i-- {
		newNode := &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		curNode.Next = newNode
		curNode = curNode.Next
	}

	return newHead.Next
}
