# https://leetcode.com/problems/subtree-of-another-tree/
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        if (not root and subRoot) or (not subRoot and root):
            return False

        if not root and not subRoot:
            return True

        current = self.sameTree(root, subRoot)
        left = self.isSubtree(root.left, subRoot)
        right = self.isSubtree(root.right, subRoot)

        return current or left or right

    def sameTree(self, root1, root2):
        if (not root1 and root2) or (not root2 and root1):
            return False

        if not root1 and not root2:
            return True

        if root1.val != root2.val:
            return False

        areLeftEqual = self.sameTree(root1.left, root2.left)
        areRightEqual = self.sameTree(root1.right, root2.right)

        return areLeftEqual and areRightEqual
