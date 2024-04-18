# https://leetcode.com/problems/evaluate-reverse-polish-notation/description/


class Solution:
    def evalRPN(self, tokens):
        stack = []

        for t in tokens:
            if self.isOperator(t):
                second = stack.pop()
                first = stack.pop()
                result = self.performOperation(first, second, t)
                stack.append(result)
            else:
                stack.append(t)

        return int(stack.pop())

    def isOperator(self, token):
        return token in ["+", "-", "*", "/"]

    def performOperation(self, first, second, operator):
        if operator == "+":
            return int(first) + int(second)
        if operator == "-":
            return int(first) - int(second)
        if operator == "*":
            return int(first) * int(second)
        if operator == "/":
            return int(first) / int(second)


s = Solution()

tokens = ["2", "1", "+", "3", "*"]
print("--> ", s.evalRPN(tokens))

tokens = ["4", "13", "5", "/", "+"]
print("--> ", s.evalRPN(tokens))

tokens = ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
print("--> ", s.evalRPN(tokens))
