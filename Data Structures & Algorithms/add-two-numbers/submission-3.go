/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    cur1 := l1
	cur2 := l2
	tmp := 0
	head := &ListNode{}
	result := head
	for cur1 != nil || cur2 != nil {
		dummy := &ListNode{}
		value := tmp
		if cur1 != nil {
			value += cur1.Val
		}
		if cur2 != nil {
			value += cur2.Val
		}

		if value > 9 {
			tmp = value / 10
			dummy.Val = value % 10
		} else {
			tmp = 0
			dummy.Val = value
		}
		if cur2 != nil {
			cur2 = cur2.Next
		}
		if cur1 != nil {
			cur1 = cur1.Next
		}
		result.Next = dummy
		result = result.Next
	}
	if tmp != 0 {
		result.Next = &ListNode{Val: tmp}
	}
	return head.Next
}
