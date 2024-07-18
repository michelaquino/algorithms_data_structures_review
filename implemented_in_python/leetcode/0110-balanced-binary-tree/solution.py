# https://leetcode.com/problems/balanced-binary-tree/


# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        depth, balance = self.checkBalance(root)
        return balance

    def checkBalance(self, root):
        if not root:
            return (0, True)

        dephtRight, isRightBalanced = self.checkBalance(root.right)
        dephtLeft, isLeftBalanced = self.checkBalance(root.left)

        depth = 1 + max(dephtRight, dephtLeft)
        isBalance = (
            abs(dephtRight - dephtLeft) <= 1 and isLeftBalanced and isRightBalanced
        )
        return (depth, isBalance)
