class Solution:
    def findMin(self, nums):
        left, right = 0, len(nums) - 1

        while left <= right:
            middle = left + ((right - left) // 2)

            if nums[middle] < nums[middle - 1]:
                return nums[middle]
            elif nums[middle] > nums[(-len(nums) + middle + 1)]:
                return nums[(-len(nums) + middle + 1)]
            else:
                if nums[0] < nums[middle] > nums[-1]:
                    left = middle + 1
                else:
                    right = middle - 1

        return None


s = Solution()
nums = [3, 4, 5, 1, 2]
print("--> ", s.findMin(nums))

nums = [4, 5, 6, 7, 0, 1, 2]
print("--> ", s.findMin(nums))

nums = [11, 13, 15, 17]
print("--> ", s.findMin(nums))
