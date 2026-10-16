# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        prev=None
        cur=head
        counter=0
        while cur:
            next=cur.next
            cur.next=prev
            prev=cur
            cur=next
            counter+=1
        if n > counter:
            return head
        reverse=prev
        if n <= counter:
            if n == 1:
                reverse=reverse.next
            else:
                cur=prev
                for i in range(n-2):
                    cur=cur.next
                cur.next=cur.next.next
        prev=None
        cur=reverse
        while cur:
            next=cur.next
            cur.next=prev
            prev=cur
            cur=next
        return prev
        
        