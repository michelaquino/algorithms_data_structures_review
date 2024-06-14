# https://leetcode.com/problems/minimum-window-substring/


class SimplerSolution:
    def minWindow(self, s: str, t: str) -> str:
        tCount = {}
        for l in t:
            tCount[l] = tCount.get(l, 0) + 1

        start, end = 0, float("infinity")
        sCount = {}
        left = 0
        for right in range(0, len(s)):
            sCount[s[right]] = sCount.get(s[right], 0) + 1

            while self.conditionIsSatisfied(tCount, sCount):
                if (right - left + 1) < (end - start + 1):
                    end = right
                    start = left

                sCount[s[left]] -= 1
                left += 1

        return s[start : end + 1] if end < float("infinity") else ""

    def conditionIsSatisfied(self, tCount, sCount):
        for letter, count in tCount.items():
            countInS = sCount.get(letter, 0)
            if countInS < count:
                return False

        return True


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if len(t) > len(s):
            return ""

        tCount = {}
        for letter in t:
            tCount[letter] = tCount.get(letter, 0) + 1

        start, end = 0, len(s) + 1

        lettersUsed = 0
        left, right = 0, 0
        while right <= len(s):
            rightLetter = s[right] if right < len(s) else s[right - 1]
            # found a substring
            if lettersUsed == len(t):
                if right - left < end - start:
                    start = left
                    end = right

                tCount[s[left]] += 1
                if tCount[s[left]] > 0:
                    lettersUsed -= 1

                left += 1
                continue

            if right < len(s):
                tCount[s[right]] = tCount.get(s[right], 0) - 1
                if tCount[s[right]] >= 0:
                    lettersUsed += 1

            right += 1

        if end > len(s):
            return ""

        return s[start:end]


class NeetcodeSolution:
    def minWindow(self, s: str, t: str) -> str:
        if t == "":
            return ""

        countT, window = {}, {}
        for c in t:
            countT[c] = 1 + countT.get(c, 0)

        have, need = 0, len(countT)
        res, resLen = [-1, -1], float("infinity")
        l = 0
        for r in range(len(s)):
            c = s[r]
            window[c] = 1 + window.get(c, 0)

            if c in countT and window[c] == countT[c]:
                have += 1

            while have == need:
                # update our result
                if (r - l + 1) < resLen:
                    res = [l, r]
                    resLen = r - l + 1
                # pop from the left of our window
                window[s[l]] -= 1
                if s[l] in countT and window[s[l]] < countT[s[l]]:
                    have -= 1
                l += 1
        l, r = res
        return s[l : r + 1] if resLen != float("infinity") else ""


sol = Solution()

s = "ADOBECODEBANC"
t = "ABC"
print(sol.minWindow(s, t))

s = "a"
t = "a"
print(sol.minWindow(s, t))

s = "a"
t = "aa"
print(sol.minWindow(s, t))

s = "a"
t = "b"
print(sol.minWindow(s, t))
