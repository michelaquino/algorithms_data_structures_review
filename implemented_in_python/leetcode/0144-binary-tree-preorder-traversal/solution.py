# https://leetcode.com/problems/binary-tree-preorder-traversal/
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
#
# O(N) time
# O(N) memory
class Solution:
    def preorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        result = []

        def dfs(root):
            if not root:
                return

            result.append(root.val)
            dfs(root.left)
            dfs(root.right)

        dfs(root)
        return result
