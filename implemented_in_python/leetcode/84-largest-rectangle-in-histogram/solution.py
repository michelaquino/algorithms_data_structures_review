# https://leetcode.com/problems/largest-rectangle-in-histogram/


class Solution:
    def largestRectangleArea(self, heights):
        # enquanto top[0] for maior que height[i]
        # remove da stack, calcula área e guarda a posição
        # Adiciona (height[i], pos) na stack
        #
        # itera na stack calculando a área dos elementos que sobraram

        largestArea = -1
        stack = []  # (index, height)
        for i, h in enumerate(heights):
            startIndex = i
            while len(stack) > 0 and stack[-1][1] > h:
                index, height = stack.pop()
                largestArea = max(largestArea, (i - index) * height)
                startIndex = height

            stack.append((startIndex, h))

        for index, height in stack:
            largestArea = max(largestArea, (len(heights) - index) * height)

        return largestArea


s = Solution()

heights = [2, 1, 5, 6, 2, 3]
print("--> ", s.largestRectangleArea(heights))


heights = [2, 4]
print("--> ", s.largestRectangleArea(heights))


heights = [2, 1, 2]
print("--> ", s.largestRectangleArea(heights))
