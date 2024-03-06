# https://leetcode.com/problems/product-of-array-except-self


def solution(nums):

    pre = [1] * len(nums)
    pos = [1] * len(nums)

    for i in range(1, len(nums)):
        pre[i] = pre[i - 1] * nums[i - 1]

    for i in range(len(nums) - 2, -1, -1):
        pos[i] = pos[i + 1] * nums[i + 1]

    result = []
    for i in range(0, len(pre)):
        result.append(pre[i] * pos[i])

    return result


print("===>", solution([1, 2, 3, 4]))
print("===>", solution([-1, 1, 0, -3, 3]))
