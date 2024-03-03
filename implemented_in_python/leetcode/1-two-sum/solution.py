def solution(nums, target):
    nums_index_hash = {}

    for i, num in enumerate(nums):

        n = target - num
        indexes = nums_index_hash.get(n, set([]))

        if len(indexes) != 0:
            another_indexes = indexes - set([i])
            if len(another_indexes) != 0:
                return [i, list(another_indexes)[0]]

        if num in nums_index_hash:
            nums_index_hash[num].add(i)
        else:
            nums_index_hash[num] = set([i])

    return []


print(solution([2, 7, 11, 15], 9))
print(solution([3, 2, 4], 6))
print(solution([3, 3], 6))
print(solution([], 6))
print(solution([3, 3], 5))
