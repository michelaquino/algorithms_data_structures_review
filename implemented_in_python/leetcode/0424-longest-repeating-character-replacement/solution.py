# https://leetcode.com/problems/longest-repeating-character-replacement/


class Solution:

    # First solution: O(n) time
    def first_characterReplacement(self, s: str, k: int) -> int:
        if len(s) == 0:
            return 0

        first, second = 0, 0
        letterMap = {}
        longest = 0

        while second < len(s):
            letterMap[s[second]] = letterMap.get(s[second], 0) + 1
            segmentSize = second - first + 1

            if segmentSize - max(letterMap.values()) > k:
                letterMap[s[first]] = letterMap.get(s[first], 0) - 1
                first += 1

            longest = max(longest, segmentSize)
            second += 1

        return longest

    # Optimized solution O(n)
    def characterReplacement(self, s: str, k: int) -> int:
        left = 0
        count = {}
        longest = 0

        for right in range(len(s)):
            count[s[right]] = count.get(s[right], 0) + 1

            while (right - left + 1) - max(count.values()) > k:
                count[s[left]] = count[s[left]] - 1
                left += 1

            longest = max(longest, right - left + 1)

        return longest





sol = Solution()
s = "ABAB"
k = 2
print(sol.characterReplacement(s, k))

s = "AABABBA"
k = 1
print(sol.characterReplacement(s, k))

s = "ABBBA"
k = 2
print(sol.characterReplacement(s, k))

s = "ABAA"
k = 0 
print(sol.characterReplacement(s, k))
