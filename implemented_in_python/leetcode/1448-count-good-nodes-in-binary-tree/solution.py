# https://leetcode.com/problems/count-good-nodes-in-binary-tree/
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def goodNodes(self, root: TreeNode) -> int:

        def dfs(root, maxPrev):
            if not root:
                return 0

            good = 0 if maxPrev > root.val else 1
            maxPrev = max(maxPrev, root.val)
            goodLeft = dfs(root.left, maxPrev)
            goodRight = dfs(root.right, maxPrev)

            return good + goodLeft + goodRight

        return dfs(root, root.val if root else 0)
