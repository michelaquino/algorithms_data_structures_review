# https://leetcode.com/problems/longest-palindromic-substring
# O(nˆ2) time
# O(1) memory
class Solution:
    def longestPalindrome(self, s: str) -> str:
        answer = ""

        def lengthOfPalindrome(left, right):
            nonlocal answer
            while left >= 0 and right < len(s) and s[left] == s[right]:
                if (right - left + 1) > len(answer):
                    answer = s[left : right + 1]

                left -= 1
                right += 1

        for i in range(len(s)):
            # odd
            lengthOfPalindrome(i, i)

            # even
            lengthOfPalindrome(i, i + 1)

        return answer
