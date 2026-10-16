# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        slow=head
        fast=head

        while fast.next and fast.next.next:
            slow=slow.next
            fast=fast.next.next

        cur=slow.next
        slow.next=None

        prev=None
        while cur:
            next=cur.next
            cur.next=prev
            prev=cur
            cur=next
        s1=head
        s2=prev
        result=ListNode()
        merged=result

        while s2:
            merged.next=s1
            s1=s1.next
            merged=merged.next
            merged.next=s2
            s2=s2.next
            merged=merged.next
        if s1:
            merged.next=s1
             
        
            