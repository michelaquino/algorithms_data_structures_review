# https://leetcode.com/problems/valid-parentheses/

# O(n) time; O(n) memory


class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack = []
        for c in s:
            if c in ["(", "{", "["]:
                stack.append(c)
                continue

            if len(stack) == 0:
                return False

            top = stack[-1]

            if (
                (c == ")" and top != "(")
                or (c == "}" and top != "{")
                or (c == "]" and top != "[")
            ):
                return False

            stack.pop()

        return len(stack) == 0


sol = Solution()
print(sol.isValid("()"))
print(sol.isValid("()[]{}"))
print(sol.isValid("(]"))
print(sol.isValid("([{{{}}}])"))
print(sol.isValid("([{{{{}}}])"))
print(sol.isValid("["))
