# https://leetcode.com/problems/word-break


# O(nˆ2 * m) time
# O(n) memory
class TopDownSolution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        memo = {}

        def dfs(startIndex):
            if startIndex >= len(s):
                return True

            if startIndex in memo:
                return memo[startIndex]

            found = False
            for word in wordDict:
                endIndex = startIndex + len(word)
                substring = s[startIndex:endIndex]

                if word == substring:
                    memo[word] = True
                    found = found or dfs(endIndex)

            memo[startIndex] = found
            return found

        return dfs(0)


# O(nˆ2 * m) time
# O(n) memory
class BottomUpSolution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        quantity = len(s) + 1
        memo = [False] * quantity
        memo[len(s)] = True

        for i in range(len(s) - 1, -1, -1):  # N
            for word in wordDict:  # M
                if i + len(word) <= len(s) and s[i : i + len(word)] == word:  # N
                    memo[i] = memo[i + len(word)]

                if memo[i]:
                    break

        return memo[0]
