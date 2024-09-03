# https://leetcode.com/problems/longest-increasing-subsequence


class Solution:
    # O(n^2) time
    # O(n) memory
    def lengthOfLIS(self, nums: List[int]) -> int:
        dp = [1] * len(nums)

        for i in range(len(nums) - 1, -1, -1):
            for j in range(i + 1, len(nums)):
                if nums[i] < nums[j]:
                    dp[i] = max(dp[i], 1 + dp[j])

        return max(dp)

    # O(n^2) time
    # O(n^2) memory
    def lengthOfLISMemoization(self, nums: List[int]) -> int:
        memo = {}

        # Memoization
        def dfs(index, last):
            if index >= len(nums):
                return 1

            if nums[index] <= last:
                return 0

            if index in memo:
                return memo[index]

            maxSequence = 0
            for i in range(index + 1, len(nums)):
                sequence = dfs(i, nums[index])
                maxSequence = max(maxSequence, sequence)

            memo[index] = 1 + maxSequence
            return 1 + maxSequence

        maxSequence = 0
        for i in range(len(nums)):
            maxSequence = max(maxSequence, dfs(i, float("-infinity")))

        return maxSequence
