# https://neetcode.io/problems/count-connected-components
# https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph


# O(n + e)
# n: vertices
# e: edges
class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        adjacency = {i: [] for i in range(n)}

        for vertice, neighboor in edges:
            adjacency[vertice].append(neighboor)
            adjacency[neighboor].append(vertice)

        visited = set()

        def dfs(vertice):
            if vertice in visited:
                return

            visited.add(vertice)
            for v in adjacency[vertice]:
                dfs(v)

        count = 0
        for vertice in range(n):
            if vertice not in visited:
                count += 1
                dfs(vertice)

        return count
