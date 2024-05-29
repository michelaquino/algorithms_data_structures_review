class Solution(object):
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """

        left, right, maximum = 0, len(height) - 1, -1

        while left < right:
            lessHeight = min(height[left], height[right])
            area = self.calcArea(lessHeight, left, right)

            maximum = max(maximum, area)

            if height[left] < height[right]:
                left += 1
            else:
                right -= 1

        return maximum

    def calcArea(self, height, left, right):
        return height * (right - left)


s = Solution()
height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
print(s.maxArea(height))

height = [1, 1]
print(s.maxArea(height))

height = [1, 2, 4, 3]
print(s.maxArea(height))
