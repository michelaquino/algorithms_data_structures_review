# https://leetcode.com/problems/sliding-window-maximum/

from collections import deque


class Solution:
    def maxSlidingWindow(self, nums, k):
        answer = []
        d = deque([])

        l = 0
        for r in range(0, len(nums)):
            # out of window
            if r - l + 1 > k:
                # print(d)
                answer.append(d[0])
                if d[0] == nums[l]:
                    d.popleft()

                l += 1

            while len(d) > 0 and d[-1] < nums[r]:
                d.pop()

            d.append(nums[r])

        answer.append(d[0])
        return answer

    # wrong solution
    def maxSlidingWindow_wrong(self, nums, k):
        maxNumber = float("-infinity")

        for i in range(0, k):
            maxNumber = max(maxNumber, nums[i])

        maxWindow = [maxNumber]
        for i in range(k, len(nums)):
            maxNumber = max(maxNumber, nums[i])
            maxWindow.append(maxNumber)

        return maxWindow


s = Solution()

nums = [1, 3, -1, -3, 5, 3, 6, 7]
k = 3
print(s.maxSlidingWindow(nums, k))

nums = [1]
k = 1
print(s.maxSlidingWindow(nums, k))

nums = [5, 4, 3, 2, 1, 0]
k = 3
print(s.maxSlidingWindow(nums, k))

nums = [9, 10, 9, -7, -4, -8, 2, -6]
k = 5
print(s.maxSlidingWindow(nums, k))
