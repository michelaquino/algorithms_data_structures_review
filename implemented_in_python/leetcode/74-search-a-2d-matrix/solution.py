class Solution:
    def searchMatrix(self, matrix, target):
        if len(matrix) == 0:
            return False

        m, n = len(matrix), len(matrix[0])

        start, end = (0, (m * n) - 1)

        while start <= end:
            middle = int((end + start) / 2)

            row, col = self.rowColByIndex(n, middle)
            number = matrix[row][col]

            if number < target:
                start = middle + 1
            elif number > target:
                end = middle - 1
            else:
                return True

        return False

    def rowColByIndex(self, n, index):
        row = index // n
        col = index - (n * row)

        return (int(row), int(col))

    # Performing two binary search
    def neetCodeSolution(self, matrix, target):
        ROWS, COLS = len(matrix), len(matrix[0])

        top, bot = 0, ROWS - 1
        while top <= bot:
            row = (top + bot) // 2
            if target > matrix[row][-1]:
                top = row + 1
            elif target < matrix[row][0]:
                bot = row - 1
            else:
                break

        if not (top <= bot):
            return False
        row = (top + bot) // 2
        l, r = 0, COLS - 1
        while l <= r:
            m = (l + r) // 2
            if target > matrix[row][m]:
                l = m + 1
            elif target < matrix[row][m]:
                r = m - 1
            else:
                return True
        return False


s = Solution()

matrix = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
target = 3
print(s.searchMatrix(matrix, target))

matrix = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
target = 13
print(s.searchMatrix(matrix, target))

# print(s.rowColByIndex(4, 0))
# print(s.rowColByIndex(4, 1))
# print(s.rowColByIndex(4, 2))
# print(s.rowColByIndex(4, 3))
# print()
# print(s.rowColByIndex(4, 4))
# print(s.rowColByIndex(4, 5))
# print(s.rowColByIndex(4, 6))
# print(s.rowColByIndex(4, 7))
# print()
# print(s.rowColByIndex(4, 8))
# print(s.rowColByIndex(4, 9))
# print(s.rowColByIndex(4, 10))
# print(s.rowColByIndex(4, 11))
