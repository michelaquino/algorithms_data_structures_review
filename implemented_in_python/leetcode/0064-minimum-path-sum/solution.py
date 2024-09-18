# https://leetcode.com/problems/minimum-path-sum
# O(N * M) time
# O(N * M) memory
class Solution_DFS_with_memo:
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


# DP
class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        default_value = float("infinity")

        for row in range(len(grid)):
            for col in range(len(grid[0])):
                if row == 0 and col == 0:
                    continue

                top = float("infinity")
                left = float("infinity")
                if row > 0:
                    top = grid[row - 1][col]

                if col > 0:
                    left = grid[row][col - 1]

                grid[row][col] += min(top, left)

        return grid[len(grid) - 1][len(grid[0]) - 1]
