# https://leetcode.com/problems/subsets/


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result = []
        subset = []

        def backtracking(index):
            if index >= len(nums):
                result.append(subset.copy())
                return

            # decision to include nums[i]
            subset.append(nums[index])
            backtracking(index + 1)

            # decision NOT to include nums[i]
            subset.pop()
            backtracking(index + 1)

        backtracking(0)
        return result
