# https://leetcode.com/problems/combination-sum-ii/
# Time: O(n * 2^n)
# Memory: O(n * 2^n)
class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()  # n log n

        result = []

        def backtracking(index, sum, subset):
            if sum == target:
                result.append(subset.copy())  # n
                return

            if sum > target or index >= len(candidates):
                return

            # use candidates[index]
            subset.append(candidates[index])
            backtracking(index + 1, sum + candidates[index], subset)

            # do not use candidates[index]
            subset.pop()
            while (
                index + 1 < len(candidates)
                and candidates[index] == candidates[index + 1]
            ):
                index += 1

            backtracking(index + 1, sum, subset)

        backtracking(0, 0, [])
        return result
