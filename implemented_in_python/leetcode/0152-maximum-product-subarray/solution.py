# https://leetcode.com/problems/maximum-product-subarray
#
# O(n) time
# O(1) memory
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        res = max(nums)
        currMin, currMax = 1, 1

        for n in nums:
            temp = currMax * n
            currMax = max(currMax * n, currMin * n, n)
            currMin = min(temp, currMin * n, n)

            res = max(res, currMax)

        return res
