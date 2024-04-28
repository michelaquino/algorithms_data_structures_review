# https://leetcode.com/problems/binary-search


class Solution:
    def search(self, nums, target):
        start, end = 0, len(nums) - 1
        while start <= end:
            # middle = (end + start) // 2 # Can have overflow if end and start are too big, for example, close to 2^31
            middle = start + ((end - start) // 2)
            num = nums[middle]

            if num < target:
                start = middle + 1
            elif num > target:
                end = middle - 1
            else:
                return middle

        return -1


s = Solution()
nums = [-1, 0, 3, 5, 9, 12]
target = 9
print("--> ", s.search(nums, target))

nums = [-1, 0, 3, 5, 9, 12]
target = 2
print("--> ", s.search(nums, target))
