# https://leetcode.com/problems/find-the-duplicate-number/
#
# https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare


class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        tortoise, hare = 0, 0
        while True:
            tortoise = nums[tortoise]
            hare = nums[nums[hare]]
            if tortoise == hare:
                break

        tortoise2 = 0
        while True:
            tortoise = nums[tortoise]
            tortoise2 = nums[tortoise2]
            if tortoise == tortoise2:
                return tortoise


class NaiveSolution:
    def findDuplicate(self, nums: List[int]) -> int:
        numSet = set()

        for n in nums:
            if n in numSet:
                return n

            numSet.add(n)
