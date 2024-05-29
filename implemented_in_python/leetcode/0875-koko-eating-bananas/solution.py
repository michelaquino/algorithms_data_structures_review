# https://leetcode.com/problems/koko-eating-bananas
import math


class Solution:
    # O(nˆ2)
    def minEatingSpeedNaive(self, piles, h):
        piles.sort()  # n log (n)

        for i in range(1, piles[-1] + 1):
            hours = self.calcHours(piles, i)

            if hours == h:
                return i

        return 0

    def minEatingSpeed(self, piles, h):
        start, end = 1, max(piles)
        minimumSoFar = end

        while start <= end:
            middle = start + ((end - start) // 2)
            hours = self.calcHours(piles, middle)
            if hours > h:
                start = middle + 1
            elif hours <= h:
                minimumSoFar = min(minimumSoFar, middle)
                end = middle - 1

        return minimumSoFar

    def calcHours(self, piles, k):
        hours = 0
        for bananas in piles:
            hours += math.ceil(bananas / k)

        return hours


s = Solution()
print("--")
piles = [3, 6, 7, 11]
h = 8
# print(s.minEatingSpeedNaive(piles, h))
print(s.minEatingSpeed(piles, h))

print("--")
piles = [30, 11, 23, 4, 20]
h = 5
# print(s.minEatingSpeedNaive(piles, h))
print(s.minEatingSpeed(piles, h))

print("--")
piles = [30, 11, 23, 4, 20]
h = 6
# print(s.minEatingSpeedNaive(piles, h))
print(s.minEatingSpeed(piles, h))

print("--")
piles = [312884470]
h = 312884469
print(s.minEatingSpeed(piles, h))
