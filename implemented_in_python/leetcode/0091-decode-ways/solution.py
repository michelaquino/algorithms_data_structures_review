# https://leetcode.com/problems/decode-ways/


# O(n) time
# O(n) memory
class DFSSolution:
    def numDecodings(self, s: str) -> int:
        memo = {}

        def dfs(index):
            if index >= len(s):
                return 1

            if index in memo:
                return memo[index]

            if s[index] == "0":
                return 0

            count = dfs(index + 1)
            if index + 1 < len(s) and int(s[index : index + 2]) <= 26:
                count += dfs(index + 2)

            memo[index] = count
            return count

        return dfs(0)


# O(n) time
# O(n) memory
class DynamicSolution:
    def numDecodings(self, s: str) -> int:
        memo = {len(s): 1}

        for index in range(len(s) - 1, -1, -1):
            if s[index] == "0":
                memo[index] = 0
            else:
                memo[index] = memo[index + 1]

            if (
                index + 1 < len(s)
                and s[index] != "0"
                and int(s[index : index + 2]) <= 26
            ):
                memo[index] += memo[index + 2]

        return memo[0]
