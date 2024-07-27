# https://leetcode.com/problems/reverse-nodes-in-k-group/


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    # O(n) time complexity, O(1) memory complexity
    def reverseKGroup(self, head, k):
        if not head:
            return head

        if not self.enoughNodes(head, k):
            return head

        curr, tail = head, head
        i = 1
        prev = None
        while curr and i <= k:
            next = curr.next
            curr.next = prev
            prev = curr
            curr = next
            i += 1

        head = prev
        tail.next = self.reverseKGroup(curr, k)
        return head

    def enoughNodes(self, curr, k):
        while curr and k > 0:
            curr = curr.next
            k -= 1

        return k == 0

    # O(n) time complexity, O(n) memory complexity
    def reverseKGroupExtraMemory(self, head, k):
        if not head:
            return head

        curr, tail = head, None

        i = 1
        prev = None
        while curr and i <= k:
            next = curr.next

            currNew = ListNode(curr.val, prev)
            if not tail:
                tail = currNew

            prev = currNew
            curr = next
            i += 1

        if i <= k:
            return head

        head = prev
        tail.next = self.reverseKGroup(curr, k)
        return head


def printList(head):
    curr = head

    while curr:
        print(curr.val, " -> ", end=" ")
        curr = curr.next

    print("")


# n6 = ListNode(6)
# n5 = ListNode(5, n6)
n5 = ListNode(5)
n4 = ListNode(4, n5)
n3 = ListNode(3, n4)
n2 = ListNode(2, n3)
n1 = ListNode(1, n2)

s = Solution()
# head = s.reverseKGroup(n1, 2)
# printList(head)

head = s.reverseKGroup(n1, 3)
printList(head)
