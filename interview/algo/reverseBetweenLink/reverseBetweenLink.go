package reverseBetweenLink

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n-m <= 1 {
		return head
	}
	// 反转后要衔接的节点
	var nodeAfterConcat *ListNode
	// 先找到要反转的头节点
	var rHead *ListNode
	i := 1
	for head != nil {
		if m == i {
			rHead = head
			break
		}
		nodeAfterConcat = head
		head = head.Next
	}
	i = m
	var newHead *ListNode
	var tmpNext *ListNode
	var afterNode *ListNode

	for rHead != nil {
		if i > n {
			afterNode = rHead
			break
		}
		tmpNext = rHead.Next
		rHead.Next = newHead
		newHead = rHead
		rHead = tmpNext
		i++
	}
	if nodeAfterConcat != nil {
		nodeAfterConcat.Next = rHead
	}
	if rHead != nil {
		rHead.Next = afterNode
	}

	return nil
}
