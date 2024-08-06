# https://leetcode.com/problems/kth-smallest-element-in-a-bst/


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        index = 0

        def dfs(root):
            nonlocal index
            if not root:
                return None

            answer = dfs(root.left)
            if answer is not None:
                return answer

            index += 1
            if index == k:
                return root.val

            answer = dfs(root.right)
            if answer is not None:
                return answer

            return None

        return dfs(root)
