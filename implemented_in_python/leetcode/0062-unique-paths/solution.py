# https://leetcode.com/problems/unique-paths/
# O(m * n) time
# O(m * n) memory
class Solution:
    def uniquePaths(self, m: int, n: int) -> int:

        field = [[1 for _ in range(n)] for _ in range(m)]

        for row in range(m):
            for col in range(n):
                if row == 0 and col == 0:
                    continue

                top = field[row - 1][col] if row > 0 else 0
                left = field[row][col - 1] if col > 0 else 0
                field[row][col] = top + left

        return field[m - 1][n - 1]


# O(m * n) time
# O(n) memory ( because update the row )
class NeetcodeSolution:
    def uniquePaths(self, m: int, n: int) -> int:
        row = [1] * n

        for i in range(m - 1):
            newRow = [1] * n
            for j in range(n - 2, -1, -1):
                newRow[j] = newRow[j + 1] + row[j]
            row = newRow
        return row[0]
