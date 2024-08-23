# https://leetcode.com/problems/pacific-atlantic-water-flow


# Optimal solution: O(N * M)
class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        ROWS, COLS = len(heights), len(heights[0])

        def dfs(row, col, visited):
            if (row, col) in visited:
                return

            visited.add((row, col))
            neigbors = [[-1, 0], [1, 0], [0, -1], [0, 1]]  # top, bottom, left, right
            for r, c in neigbors:
                newRow = row + r
                newCol = col + c

                # out of bounds
                if newRow < 0 or newRow >= ROWS or newCol < 0 or newCol >= COLS:
                    continue

                if heights[newRow][newCol] < heights[row][col]:
                    continue

                dfs(newRow, newCol, visited)

        reachPacific = set()
        reachAtlantic = set()
        for row in range(ROWS):
            dfs(row, 0, reachPacific)
            dfs(row, COLS - 1, reachAtlantic)

        for col in range(COLS):
            dfs(0, col, reachPacific)
            dfs(ROWS - 1, col, reachAtlantic)

        result = []
        for row in range(ROWS):
            for col in range(COLS):
                if (row, col) in reachPacific and (row, col) in reachAtlantic:
                    result.append((row, col))

        return result


# O((N * M)ˆ2)
class NaiveSolution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:

        def shouldAdd(coordinate, row, col):
            if coordinate[0] < 0 or coordinate[0] >= len(heights):
                return True

            if coordinate[1] < 0 or coordinate[1] >= len(heights[0]):
                return True

            r, c = coordinate[0], coordinate[1]
            if heights[r][c] <= heights[row][col]:
                return True

            return False

        def reachOceans(row, col):
            reachPacific = False
            reachAtlantic = False

            visited = set()
            queue = [(row, col)]  # (row, col)
            while queue:
                r, c = queue.pop()
                if (r, c) in visited:
                    continue

                visited.add((r, c))

                if r < 0 or c < 0:
                    reachPacific = True
                    continue

                if r >= len(heights) or c >= len(heights[0]):
                    reachAtlantic = True
                    continue

                top = (r - 1, c)
                if shouldAdd(top, r, c):
                    queue.append(top)

                bottom = (r + 1, c)
                if shouldAdd(bottom, r, c):
                    queue.append(bottom)

                left = (r, c - 1)
                if shouldAdd(left, r, c):
                    queue.append(left)

                right = (r, c + 1)
                if shouldAdd(right, r, c):
                    queue.append(right)

            return reachPacific and reachAtlantic

        result = []
        for row in range(len(heights)):
            for col in range(len(heights[0])):
                if reachOceans(row, col):
                    result.append([row, col])

        return result
