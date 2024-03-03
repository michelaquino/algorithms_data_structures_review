# https://leetcode.com/problems/group-anagrams
import collections


# s * l log(l)
# s = quantity of string
# l = average size of strings
def solution(strs):
    groups = {}

    for str in strs:
        str_sorted = sort_str(str)
        if str_sorted in groups:
            groups[str_sorted].append(str)
        else:
            groups[str_sorted] = [str]

    result = []
    for i in groups.values():
        result.append(i)

    return result


def sort_str(str):
    arr_sorted = sorted(str)
    return "".join(arr_sorted)


# print(solution(["eat", "tea", "tan", "ate", "nat", "bat"]))
# print(solution([""]))
# print(solution(["a"]))


def better_solution(strs):
    answer = collections.defaultdict(list)

    for s in strs:
        count = [0] * 26

        for letter in s:
            count[ord(letter) - ord("a")] += 1

        answer[tuple(count)].append(s)

    return answer.values()


print(better_solution(["eat", "tea", "tan", "ate", "nat", "bat"]))
print(better_solution([""]))
print(better_solution(["a"]))
