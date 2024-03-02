# https://leetcode.com/problems/contains-duplicate/


def solution(nums):
    set_numbers = set({})

    for n in nums:
        if n in set_numbers:
            return True

        set_numbers.add(n)

    return False


nums1 = [1, 2, 3, 1]
print(solution(nums1))

nums2 = [1, 2, 3, 4]
print(solution(nums2))

nums3 = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
print(solution(nums3))
