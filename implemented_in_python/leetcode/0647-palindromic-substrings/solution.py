# https://leetcode.com/problems/palindromic-substrings
#
# O(n^2) time
# O(1) memory


class Solution:
    def countSubstrings(self, s: str) -> int:
        total = 0
        for i in range(len(s)):
            # odd
            total += self.palindromeSubstrings(s, i, i)

            # even
            total += self.palindromeSubstrings(s, i, i + 1)

        return total

    def palindromeSubstrings(self, s, start, end):
        count = 0
        while start >= 0 and end < len(s) and s[start] == s[end]:
            count += 1
            start -= 1
            end += 1

        return count
