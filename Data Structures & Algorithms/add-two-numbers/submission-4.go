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
	carry := 0
	head := &ListNode{}
	cur := head
	for cur1 != nil || cur2 != nil || carry != 0 {
		sum := carry
		if cur1 != nil {
			sum += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			sum += cur2.Val
			cur2 = cur2.Next
		}

	
		cur.Next = &ListNode{Val: sum % 10}
		carry = sum / 10
		cur = cur.Next
	}

	return head.Next
}
