# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        slow=fast=head

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

        while s2:
            next1 = s1.next
            next2 = s2.next
            s1.next=s2
            s2.next=next1
            s1=next1
            s2=next2
             
        
            