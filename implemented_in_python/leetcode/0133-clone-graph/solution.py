"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

# https://leetcode.com/problems/clone-graph

from typing import Optional


# Using BFS
class Solution:
    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        nodeReference = {}

        visited = set()
        queue = [node]
        while queue:
            n = queue.pop()
            if not n or n.val in visited:
                continue

            visited.add(n.val)

            if n.val in nodeReference:
                continue

            nodeReference[n.val] = Node(n.val)
            queue = queue + n.neighbors

        visited = set()
        queue = [node]
        while queue:
            n = queue.pop()
            if not n or n.val in visited:
                continue

            visited.add(n.val)
            queue = queue + n.neighbors

            newNode = nodeReference[n.val]
            neighbors = []

            for neighbor in n.neighbors:
                neighbors.append(nodeReference[neighbor.val])

            newNode.neighbors = neighbors

        return nodeReference[node.val] if node else None


# Using DFS. Neetcode solution
class NeetcodeSolution:
    def cloneGraph(self, node: "Node") -> "Node":
        oldToNew = {}

        def dfs(node):
            if node in oldToNew:
                return oldToNew[node]

            copy = Node(node.val)
            oldToNew[node] = copy
            for nei in node.neighbors:
                copy.neighbors.append(dfs(nei))
            return copy

        return dfs(node) if node else None
