# https://leetcode.com/problems/product-of-array-except-self


# O(N) time and O(N) space
def solution1(nums):
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


# O(N) time and O(1) space
def solution2(nums):
    result = [1] * len(nums)

    for i in range(1, len(nums)):
        result[i] = nums[i - 1] * result[i - 1]

    postfix = 1
    for i in range(len(nums) - 1, -1, -1):
        result[i] *= postfix
        postfix *= nums[i]

    return result


print("===>", solution2([1, 2, 3, 4]))
print("===>", solution2([-1, 1, 0, -3, 3]))
