# https://leetcode.com/problems/minimum-path-sum
# O(N * M) time
# O(N * M) memory
class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        default_value = float("infinity")
        memo = {}

        def dfs(row, col):
            if row == len(grid) - 1 and col == len(grid[0]) - 1:
                return grid[row][col]

            if row >= len(grid) or col >= len(grid[0]):
                return float("infinity")

            if (row, col) in memo:
                return memo[(row, col)]

            right = dfs(row, col + 1)
            down = dfs(row + 1, col)

            result = grid[row][col] + min(right, down)
            memo[(row, col)] = result
            return result

        return dfs(0, 0)
