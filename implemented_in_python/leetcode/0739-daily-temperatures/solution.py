# https://leetcode.com/problems/daily-temperatures


class Solution:
    def dailyTemperatures(self, temperatures):
        answer = [0] * len(temperatures)

        stack = []

        for day in range(len(temperatures) - 1, -1, -1):
            temp = temperatures[day]
            while len(stack) > 0 and stack[-1][1] <= temp:
                stack.pop()

            if len(stack) > 0:
                answer[day] = stack[-1][0] - day

            stack.append([day, temp])

        return answer


s = Solution()

temperatures = [73, 74, 75, 71, 69, 72, 76, 73]
print("--> ", s.dailyTemperatures(temperatures))

temperatures = [30, 40, 50, 60]
print("--> ", s.dailyTemperatures(temperatures))

temperatures = [30, 60, 90]
print("--> ", s.dailyTemperatures(temperatures))

temperatures = [89, 62, 70, 58, 47, 47, 46, 76, 100, 70]
print("--> ", s.dailyTemperatures(temperatures))
