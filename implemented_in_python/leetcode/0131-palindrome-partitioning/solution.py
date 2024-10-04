# https://leetcode.com/problems/palindrome-partitioning
# Time: O(n * 2^n)
# Memory: O(n * 2^n)
class Solution:
    def partition(self, s: str) -> List[List[str]]:
        result = []

        def backtracking(index, subset):
            if index >= len(s):
                result.append(subset.copy())
                return

            for i in range(index, len(s)):
                if self.isPalindrome(s, index, i):
                    subset.append(s[index : i + 1])
                    backtracking(i + 1, subset)
                    subset.pop()

        backtracking(0, [])
        return result

    def isPalindrome(self, word, start, end):
        while start <= end:
            if word[start] != word[end]:
                return False

            start += 1
            end -= 1
        return True
