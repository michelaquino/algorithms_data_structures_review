class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """

        # n log(n)
        nums.sort()

        answer = []
        for index, number in enumerate(nums):

            if index > 0 and number == nums[index - 1]:
                continue

            left, right = index + 1, len(nums) - 1

            while left < right:
                threeSum = number + nums[left] + nums[right]

                if threeSum < 0:
                    left += 1
                elif threeSum > 0:
                    right -= 1
                else:
                    answer.append([number, nums[left], nums[right]])
                    left += 1
                    right -= 1
                    while nums[left] == nums[left - 1] and left < right:
                        left += 1

        return answer


s = Solution()
# nums = [-1, 0, 1, 2, -1, -4]
# print(s.threeSum(nums))

nums = [-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4]
print(s.threeSum(nums))
