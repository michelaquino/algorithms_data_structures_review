# https://leetcode.com/problems/binary-tree-right-side-view/
#
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    class Solution:
        def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
            answer = []

            queue = collections.deque([root])
            while len(queue) > 0:
                mostRight = None
                qLen = len(queue)

                for i in range(qLen):
                    node = queue.popleft()
                    if not node:
                        continue

                    mostRight = node
                    queue.append(node.left)
                    queue.append(node.right)

                if mostRight:
                    answer.append(mostRight.val)

            return answer
