# https://leetcode.com/problems/trapping-rain-water/


class Solution(object):

    # O(n) and O(n) extra memory
    def trap_extra_memory(self, height):
        waterTrapped = 0

        maxLeft = [0] * len(height)
        maxRight = [0] * len(height)

        for i in range(1, len(height)):
            maxLeft[i] = max(maxLeft[i - 1], height[i - 1])

        for i in range(len(height) - 2, -1, -1):
            maxRight[i] = max(maxRight[i + 1], height[i + 1])

        for i, h in enumerate(height):
            waterUnits = min(maxLeft[i], maxRight[i]) - h
            if waterUnits < 0:
                continue

            waterTrapped += waterUnits

        return waterTrapped

    # O(n) complexity and O(1) memory
    def trap(self, height):
        l, r = 0, len(height) - 1
        maxLeft, maxRight = height[l], height[r]
        result = 0

        while l < r:
            if maxLeft < maxRight:
                l += 1
                maxLeft = max(maxLeft, height[l])
                result += maxLeft - height[l]
            else:
                r -= 1
                maxRight = max(maxRight, height[r])
                result += maxRight - height[r]

        return result


s = Solution()
height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
print("6 =>", s.trap(height))

height = [4, 2, 0, 3, 2, 5]
print("9 =>", s.trap(height))

height = [4, 2, 3]
print("1 =>", s.trap(height))
