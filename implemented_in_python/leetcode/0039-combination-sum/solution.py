# https://leetcode.com/problems/combination-sum/
# O(2^t) time
# O(2^t) memory


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        subset = []

        def backtracking(i, total):
            if total > target or i >= len(candidates):
                return

            if total == target:
                result.append(subset.copy())
                return

            # decision to include current
            subset.append(candidates[i])
            backtracking(i, total + candidates[i])
            subset.pop()

            # decision not to include current
            backtracking(i + 1, total)

        backtracking(0, 0)
        return result
