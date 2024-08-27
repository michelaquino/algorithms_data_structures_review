# https://leetcode.com/problems/climbing-stairs/


# O(n) Time
# O(n) Memory
class Solution:
    def climbStairs(self, n: int) -> int:
        memo = {}

        def backtracking(step):
            if step > n:
                return 0

            if step == n:
                return 1

            if step in memo:
                return memo[step]

            result = backtracking(step + 1) + backtracking(step + 2)
            memo[step] = result
            return result

        return backtracking(0)


# O(n) Time
# O(1) Memory
class NeetcodeSolution:
    def climbStairs(self, n: int) -> int:
        if n <= 3:
            return n
        n1, n2 = 2, 3

        for i in range(4, n + 1):
            temp = n1 + n2
            n1 = n2
            n2 = temp
        return n2
