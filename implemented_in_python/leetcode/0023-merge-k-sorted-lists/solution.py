# https://leetcode.com/problems/merge-k-sorted-lists/


# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
#
# O(k * n)
# k = number of lists
class BruteForceSolution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        pointers = {}
        for i, list in enumerate(lists):
            pointers[i] = list

        dummy = ListNode()
        curr = dummy
        while curr:
            index = self.findMin(pointers)
            minNode = pointers[index] if index >= 0 else None

            curr.next = minNode
            curr = curr.next

            if minNode:
                pointers[index] = pointers[index].next

        return dummy.next

    def findMin(self, pointers):
        minIndex = -1
        minValue = float("infinity")

        for index, head in pointers.items():
            if head and head.val < minValue:
                minValue = head.val
                minIndex = index

        return minIndex


# O(n log(k))
# The main ideia is to iterate through the array of lists by two, get two lists, merge them,
# append the merged lists on a variable and update the main list with the result temporary variable merged lists.
# This way you will merge two lists at time, reducing by half the amount of lists need to be merged.
# Complexity: n log(k) , where k is the amount of lists and n is the size of the biggest list
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        while len(lists) > 1:
            mergedLists = []

            for i in range(0, len(lists), 2):
                list1 = lists[i]
                list2 = lists[i + 1] if i + 1 < len(lists) else None
                merged = self.mergeTwoLists(list1, list2)
                mergedLists.append(merged)

            lists = mergedLists

        return lists[0] if len(lists) > 0 else None

    def mergeTwoLists(self, node1, node2):
        if not node1:
            return node2

        if not node2:
            return node1

        smaller, bigger = node1, node2
        if node2.val < node1.val:
            smaller, bigger = bigger, smaller

        smaller.next = self.mergeTwoLists(smaller.next, bigger)
        return smaller
