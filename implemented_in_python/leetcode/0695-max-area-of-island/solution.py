# https://leetcode.com/problems/max-area-of-island/


# Time: O(m * n)
# Memory: O(m * n) - recursive
class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        ROWS, COLS = len(grid), len(grid[0])

        def dfs(row, col):
            # out of bounds
            if row < 0 or row >= ROWS or col < 0 or col >= COLS:
                return 0

            # water
            if grid[row][col] == 0:
                return 0

            # mark visited
            grid[row][col] = 0

            top = dfs(row - 1, col)
            down = dfs(row + 1, col)
            right = dfs(row, col - 1)
            left = dfs(row, col + 1)

            return 1 + top + down + right + left

        maxArea = 0
        for row in range(len(grid)):  # m
            for col in range(len(grid[0])):  # n
                maxArea = max(maxArea, dfs(row, col))

        return maxArea
