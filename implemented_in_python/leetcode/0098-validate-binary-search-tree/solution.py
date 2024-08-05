# https://leetcode.com/problems/validate-binary-search-tree/
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:

        def valid(root, leftBoundary, rightBoundary):
            if not root:
                return True

            if root.val <= leftBoundary or root.val >= rightBoundary:
                return False

            leftValid = valid(root.left, leftBoundary, root.val)
            rightValid = valid(root.right, root.val, rightBoundary)

            return leftValid and rightValid

        return valid(root, float("-infinity"), float("infinity"))
