# https://leetcode.com/problems/house-robber


# O(n) time
# O(1) memory
class Solution:
    def rob(self, nums: List[int]) -> int:
        rob1, rob2 = 0, 0

        for n in nums:
            temp = max(n + rob1, rob2)
            rob1 = rob2
            rob2 = temp

        return rob2


# O(n) time
# O(n) memory
class ExtraMemorySolution:
    def rob(self, nums: List[int]) -> int:
        robPerHouse = [0] * len(nums)

        for i in range(len(nums)):
            rob1 = nums[i] + robPerHouse[i - 2] if i - 2 >= 0 else nums[i]
            rob2 = robPerHouse[i - 1] if i - 1 >= 0 else nums[i]

            robPerHouse[i] = max(rob1, rob2)

        return robPerHouse[-1]
