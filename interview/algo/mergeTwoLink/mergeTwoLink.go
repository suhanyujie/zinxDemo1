package mergeTwoLink

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Merge BM4 合并两个排序的链表
 * 题目地址：https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337?tpId=295&tqId=23267&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	newHead := &ListNode{}
	curNode := newHead
	p1 := pHead1
	p2 := pHead2

	for p1 != nil && p2 != nil {
		if p1.Val >= p2.Val {
			curNode.Next = p2
			p2 = p2.Next
		} else {
			curNode.Next = p1
			p1 = p1.Next
		}
		curNode = curNode.Next
	}
	// 衔接某个链表不为空时的后续节点
	if p1 != nil {
		curNode.Next = p1
	}
	if p2 != nil {
		curNode.Next = p2
	}

	return newHead.Next
}
