# https://leetcode.com/problems/diameter-of-binary-tree/


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        def dfs(root):
            if not root:
                return (0, 0)  # (diameter, max depth)

            # calc the diameter and max depth
            diameterLeft, maxLeft = dfs(root.left)
            diameterRight, maxRight = dfs(root.right)

            # current diameter
            diameter = maxLeft + maxRight

            maxDiameter = max(diameter, max(diameterLeft, diameterRight))
            maxDepth = 1 + max(maxLeft, maxRight)

            return (maxDiameter, maxDepth)

        return dfs(root)[0]
