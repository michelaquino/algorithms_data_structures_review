# https://leetcode.com/problems/valid-anagram


def solution(s, t):
    if s == "" and t == "":
        return False

    if len(s) != len(t):
        return False

    hash_table = {}
    for l in s:
        if l in hash_table:
            hash_table[l] += 1
        else:
            hash_table[l] = 1

    for l in t:
        if l in hash_table:
            quantity = hash_table[l]
            if quantity <= 0:
                return False
            else:
                hash_table[l] -= 1
        else:
            return False

    return True


print(solution("anagram", "nagaram"))
print(solution("rat", "car"))
print(solution("", ""))
print(solution("a", ""))
