# https://leetcode.com/problems/longest-substring-without-repeating-characters/


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) < 2:
            return len(s)

        first, second = 0, 1
        letterCount = {s[first]: 1}
        longest = 0

        while second < len(s):
            count = letterCount.get(s[second], 0)
            letterCount[s[second]] = count + 1

            if count > 0:
                longest = max(longest, second - first)
                while s[first] != s[second]:
                    letterCount[s[first]] -= 1
                    first += 1

                letterCount[s[first]] -= 1
                first += 1

            second += 1

        return max(longest, second - first)

    def simpler_lengthOfLongestSubstring(self, s: str) -> int:
        letterSet = set()
        longest = 0
        first, second = 0, 0

        while second < len(s):
            while s[second] in letterSet:
                letterSet.remove(s[first])
                first += 1

            letterSet.add(s[second])
            longest = max(longest, second - first + 1)
            second += 1

        return longest


s = Solution()


print(s.lengthOfLongestSubstring("abcabcbb"))
print(s.lengthOfLongestSubstring("bbbbb"))
print(s.lengthOfLongestSubstring("pwwkew"))
