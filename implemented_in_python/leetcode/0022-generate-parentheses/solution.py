# https://leetcode.com/problems/generate-parentheses/


class Solution:
    def generateParenthesis(self, n):
        # You can add an open parentheses if open count < n
        # You can add close parentheses if close count < open count
        # If close count = open count = n => add to result

        stack = []
        ans = []

        def backtracking(openCount, closeCount):
            if openCount == closeCount == n:
                ans.append("".join(stack))
                return

            if openCount < n:
                stack.append("(")
                backtracking(openCount + 1, closeCount)
                stack.pop()

            if closeCount < openCount:
                stack.append(")")
                backtracking(openCount, closeCount + 1)
                stack.pop()

        backtracking(0, 0)
        return ans


class OtherSolution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        result = []
        stack = []

        def backtracking(startCount, endCount):
            if startCount > n or endCount > startCount:
                return

            # found a result
            if len(stack) == n * 2:
                result.append("".join(stack))
                return

            stack.append("(")
            backtracking(startCount + 1, endCount)
            stack.pop()

            stack.append(")")
            backtracking(startCount, endCount + 1)
            stack.pop()

        backtracking(0, 0)
        return result


s = Solution()

n = 3
print("n=3 ===> ", s.generateParenthesis(n))
n = 1
print("n=1 ===> ", s.generateParenthesis(n))
