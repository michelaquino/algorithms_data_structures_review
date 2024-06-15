# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        """
        Do not return anything, modify head in-place instead.
        """

        # find the middle
        middle, fast = head, head
        while fast and fast.next and fast.next.next:
            middle = middle.next
            fast = fast.next.next

        # reverse second part
        tail = middle.next
        prev = None
        curr = tail
        while curr:
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next

        tail = prev

        # new appointments, merge two lists
        middle.next = None
        curr1 = head
        curr2 = tail
        while curr1 and curr2:
            curr2Next = curr2.next
            curr1Next = curr1.next

            curr2.next = curr1.next
            curr1.next = curr2

            curr1 = curr1Next
            curr2 = curr2Next

        return head
