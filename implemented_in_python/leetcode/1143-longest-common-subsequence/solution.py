# O(t1 * t2) time
# O(t1 * t2) memory
# Top down
class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        matrix = [[0 for _ in range(len(text1))] for _ in range(len(text2))]

        for row in range(len(matrix)):
            for col in range(len(matrix[0])):

                left = matrix[row][col - 1] if col > 0 else 0
                top = matrix[row - 1][col] if row > 0 else 0
                topLeft = matrix[row - 1][col - 1] if row > 0 and col > 0 else 0

                if text1[col] == text2[row]:
                    matrix[row][col] = 1 + topLeft
                else:
                    matrix[row][col] = max(left, top)

        return matrix[len(text2) - 1][len(text1) - 1]


# Bottom up
class NeetcodeSolution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        dp = [[0 for j in range(len(text2) + 1)] for i in range(len(text1) + 1)]

        for i in range(len(text1) - 1, -1, -1):
            for j in range(len(text2) - 1, -1, -1):
                if text1[i] == text2[j]:
                    dp[i][j] = 1 + dp[i + 1][j + 1]
                else:
                    dp[i][j] = max(dp[i][j + 1], dp[i + 1][j])

        return dp[0][0]
