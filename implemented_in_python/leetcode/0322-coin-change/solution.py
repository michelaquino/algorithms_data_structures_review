# https://leetcode.com/problems/coin-change/
#
# O(amount * n)
# O(n) memory


class DFSSolution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        memo = {}

        def dfs(amount):
            if amount in memo:
                return memo[amount]

            if amount == 0:
                return 0

            if amount < 0:
                return -1

            fewestCoins = float("infinity")
            for c in coins:
                result = dfs(amount - c)
                memo[amount - c] = result

                if result >= 0:
                    fewestCoins = min(fewestCoins, result)

            return fewestCoins + 1

        result = dfs(amount)
        return result if result < float("infinity") else -1


# O(amount * n)
# O(n) memory
# Dynamic programming
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [amount + 1] * (amount + 1)
        dp[0] = 0

        for a in range(1, amount + 1):
            for c in coins:
                if a - c >= 0:
                    dp[a] = min(dp[a], 1 + dp[a - c])
        return dp[amount] if dp[amount] != amount + 1 else -1
