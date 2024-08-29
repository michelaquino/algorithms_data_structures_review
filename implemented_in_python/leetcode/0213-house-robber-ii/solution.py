# https://leetcode.com/problems/house-robber-ii/
#
# O(n) time
# O(1) memory
class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]

        return max(self.maxMoney(nums[:-1]), self.maxMoney(nums[1:]))

    def maxMoney(self, nums):
        rob1, rob2 = 0, 0

        for n in nums:
            money = max(n + rob1, rob2)
            rob1 = rob2
            rob2 = money

        return rob2
