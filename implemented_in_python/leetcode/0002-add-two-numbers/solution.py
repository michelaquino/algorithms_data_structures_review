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

    def add(self, l1, l2, carry):
        if not l1 and not l2:
            return None if carry == 0 else ListNode(carry)

        first = l1.val if l1 else 0
        second = l2.val if l2 else 0

        sum = first + second + carry
        value = sum if sum < 10 else sum % 10
        carry = 0 if sum < 10 else 1

        node = ListNode(value)

        l1Next = l1.next if l1 else None
        l2Next = l2.next if l2 else None

        node.next = self.add(l1Next, l2Next, carry)
        return node


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class IterativeSolution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        dummy = ListNode()
        cur = dummy

        carry = 0
        while l1 or l2 or carry:
            v1 = l1.val if l1 else 0
            v2 = l2.val if l2 else 0

            # new digit
            val = v1 + v2 + carry
            carry = val // 10
            val = val % 10
            cur.next = ListNode(val)

            # update ptrs
            cur = cur.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None

        return dummy.next
