# https://leetcode.com/problems/top-k-frequent-elements/


def solution(nums, k):
    count = {}
    freq = [[] for i in range(len(nums) + 1)]

    for n in nums:
        count[n] = 1 + count.get(n, 0)

    for n, c in count.items():
        freq[c].append(n)

    answer = []

    for i in range(len(freq) - 1, 0, -1):
        for n in freq[i]:
            answer.append(n)

            if len(answer) == k:
                return answer

    return answer


print(solution([1, 1, 1, 2, 2, 3], 2))
print(solution([1, 1, 1, 2, 2, 2, 3, 3, 4], 3))
print(solution([1], 1))
