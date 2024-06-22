# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

# https://leetcode.com/problems/add-two-numbers/


# O(max(l1, l2) time, O(max(l1, l2)) memory
class Solution:
    def addTwoNumbers(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        return self.add(l1, l2, 0)

    def add(self, l1, l2, toSum):
        if not l1 and not l2:
            return None if toSum == 0 else ListNode(toSum)

        first = l1.val if l1 else 0
        second = l2.val if l2 else 0

        sum = first + second + toSum
        value = sum if sum < 10 else sum % 10
        toSum = 0 if sum < 10 else 1

        node = ListNode(value)

        l1Next = l1.next if l1 else None
        l2Next = l2.next if l2 else None

        node.next = self.add(l1Next, l2Next, toSum)
        return node
