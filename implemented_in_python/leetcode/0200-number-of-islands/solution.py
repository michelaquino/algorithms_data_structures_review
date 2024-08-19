# https://leetcode.com/problems/number-of-islands/


# O(m * n) time
# O(m * n) memory
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        def bfs(startRow, startColumn):
            queue = [[startRow, startColumn]]
            landFound = False

            while queue:
                row, column = queue.pop()
                # out of bound
                if row < 0 or row >= len(grid) or column < 0 or column >= len(grid[0]):
                    continue

                # already visited
                point = grid[row][column]
                if point == ".":
                    continue

                # mark visited
                grid[row][column] = "."

                # water
                if point == "0":
                    continue

                landFound = True
                queue.append([row - 1, column])
                queue.append([row + 1, column])
                queue.append([row, column - 1])
                queue.append([row, column + 1])

            return landFound

        result = 0
        for row in range(len(grid)):
            for column in range(len(grid[row])):
                landFound = bfs(row, column)
                if landFound:
                    result += 1

        return result
