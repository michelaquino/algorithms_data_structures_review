# https://leetcode.com/problems/permutations


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        result = []
        subset = []

        def backtracking(numbers):
            if len(subset) == len(nums):
                result.append(subset.copy())
                return

            for i, n in enumerate(numbers):
                subset.append(n)
                backtracking(numbers[:i] + numbers[i + 1 :])
                subset.pop()

        backtracking(nums)
        return result
