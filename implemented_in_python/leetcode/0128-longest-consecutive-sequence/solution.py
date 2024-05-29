# https://leetcode.com/problems/longest-consecutive-sequence/


def solution(nums):
    numsSet = set(nums)
    longest = 0

    for n in nums:
        # not the start of a sequence
        if n - 1 in numsSet:
            continue

        length = 1
        while n + length in numsSet:
            length += 1

        longest = max(longest, length)

    return longest


print(solution([100, 4, 200, 1, 3, 2]))
print(solution([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]))
