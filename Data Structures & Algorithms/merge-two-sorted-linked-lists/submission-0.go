/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    result := &ListNode{}
	merged := result
	for list1 != nil || list2 != nil {
		if list1 == nil {
			merged.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
			merged = merged.Next
			continue
		}
		if list2 == nil {
			merged.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
			merged = merged.Next
			continue
		}
		if list1.Val < list2.Val {
			if merged == nil {
				merged = &ListNode{Val: list1.Val}
			}
			merged.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			if merged == nil {
				merged = &ListNode{Val: list2.Val}
			}
			merged.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		merged = merged.Next
	}

	return result.Next
}
