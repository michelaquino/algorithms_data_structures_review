# https://leetcode.com/problems/graph-valid-tree
# O(n + e) -> n: vertices e: edges
class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        adjacencyMap = {}
        for vertice, neighboor in edges:
            verticeEdges = adjacencyMap.get(vertice, [])
            verticeEdges.append(neighboor)
            adjacencyMap[vertice] = verticeEdges

            neighboorEdges = adjacencyMap.get(neighboor, [])
            neighboorEdges.append(vertice)
            adjacencyMap[neighboor] = neighboorEdges

        visited = set()

        def dfs(vertice, lastNeighboor):
            if vertice in visited:
                return False

            visited.add(vertice)
            for neighboor in adjacencyMap.get(vertice, []):
                if neighboor == lastNeighboor:
                    continue

                if not dfs(neighboor, vertice):
                    return False

            return True

        valid = dfs(0, -1)
        for vertice in range(n):
            if vertice not in visited:
                return False

        return valid
