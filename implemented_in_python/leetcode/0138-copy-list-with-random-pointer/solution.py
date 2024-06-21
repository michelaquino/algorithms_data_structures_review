"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""


class FirstAttemptSolution:
    def copyRandomList(self, head: "Optional[Node]") -> "Optional[Node]":
        # map current nodes to index
        nodeIndexMap = {}
        curr = head
        index = 0
        while curr:
            nodeIndexMap[curr] = index
            curr = curr.next
            index += 1

        # clone list without update random, store the mapping index -> node
        indexNewNodesMap = {}

        def cloneList(head, index):
            if not head:
                return head

            newNext = cloneList(head.next, index + 1)
            newHead = Node(head.val, newNext, None)
            indexNewNodesMap[index] = newHead
            return newHead

        newHead = cloneList(head, 0)

        # Update the random pointer
        curr = head
        newCurr = newHead
        while curr and newCurr:
            if curr.random:
                index = nodeIndexMap[curr.random]
                newCurr.random = indexNewNodesMap[index]

            curr = curr.next
            newCurr = newCurr.next

        return newHead


# simpler solution
class Solution:
    def copyRandomList(self, head: "Optional[Node]") -> "Optional[Node]":
        oldToCopy = {}
        curr = head
        while curr:
            copy = Node(curr.val)
            oldToCopy[curr] = copy
            curr = curr.next

        curr = head
        while curr:
            copy = oldToCopy[curr]
            copy.next = oldToCopy[curr.next] if curr.next else None
            copy.random = oldToCopy[curr.random] if curr.random else None

            curr = curr.next

        return oldToCopy[head] if head else None
