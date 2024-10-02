# https://leetcode.com/problems/subsets-ii/


# Time: O(n 2^n)
#   2^n: 2 choices in n numbers on array
#   n: largest subset
# Memory: O(n 2^n)
#   2^n: 2 choices in n numbers on array
#   n: largest subset
class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        subsets = set()
        subset = []

        def backtracking(index):
            if index >= len(nums):
                subsets.add(tuple(subset.copy()))
                return

            # add
            subset.append(nums[index])
            backtracking(index + 1)

            # not add
            subset.pop()
            backtracking(index + 1)

        backtracking(0)

        result = []
        for subset in subsets:
            result.append(subset)

        return result


# Time: O(n 2^n)
#   2^n: 2 choices in n numbers on array
#   n: largest subset
# Memory: O(n 2^n)
#   2^n: 2 choices in n numbers on array
#   n: largest subset
class NeetcodeSolution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        result = []
        subset = []

        def backtracking(index):
            if index >= len(nums):
                result.append(subset.copy())
                return

            # add the numbers of nums[i]
            subset.append(nums[index])
            backtracking(index + 1)
            subset.pop()

            # Do not add the numbers of nums[i]
            # Skipping dupplicated numbers
            while index + 1 < len(nums) and nums[index + 1] == nums[index]:
                index += 1

            backtracking(index + 1)

        backtracking(0)
        return result
