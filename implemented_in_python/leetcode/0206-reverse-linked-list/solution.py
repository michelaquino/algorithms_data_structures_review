# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        curr = head
        while curr is not None:
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next

        head = prev
        return head

    def reverseListRecursive(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None:
            return None

        newHead = head
        if head.next is not None:
            newHead = self.reverseList(head.next)
            head.next.next = head

        head.next = None
        return newHead
