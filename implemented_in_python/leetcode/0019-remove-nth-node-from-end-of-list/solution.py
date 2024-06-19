# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class SolutionPassingTwoTimes:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        size = 0
        curr = head
        while curr:
            size += 1
            curr = curr.next

        indexToRemove = size - n

        index = 0
        prev, curr = None, head
        while curr:
            if index < indexToRemove:
                prev = curr
                curr = curr.next
                index += 1
                continue

            if prev is None:
                head = curr.next
            else:
                prev.next = curr.next

            break

        return head


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        right = head
        index = 0
        while right and index < n:
            right = right.next
            index += 1

        prevLeft, left = None, head
        while right:
            prevLeft = left
            left = left.next
            right = right.next

        if prevLeft:
            prevLeft.next = left.next
        else:
            head = left.next

        return head
