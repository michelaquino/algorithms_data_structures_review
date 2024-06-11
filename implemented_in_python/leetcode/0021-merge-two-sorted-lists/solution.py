# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    # Iterative
    # T: O(l1 + l2) M: O(1)
    def mergeTwoLists(self, list1, list2):
        curr1, curr2 = list1, list2
        newHead = curr = None

        def updateCurr(nextNode):
            nonlocal curr
            nonlocal newHead

            if not curr:
                curr = nextNode
                newHead = curr
                return

            curr.next = nextNode
            curr = curr.next

        while curr1 or curr2:
            if not curr2:
                updateCurr(curr1)
                curr1 = curr1.next
            elif not curr1:
                updateCurr(curr2)
                curr2 = curr2.next
            else:
                if curr1.val < curr2.val:
                    updateCurr(curr1)
                    curr1 = curr1.next
                else:
                    updateCurr(curr2)
                    curr2 = curr2.next

        return newHead

    # Recursive
    # T: O(l1 + l2) M: O(l1 + l2)
    def mergeTwoListsRecursive(self, list1, list2):
        if not list1:
            return list2
        if not list2:
            return list1

        smaller, bigger = (list1, list2) if list1.val <= list2.val else (list2, list1)
        smaller.next = self.mergeTwoListsRecursive(smaller.next, bigger)
        return smaller
