# https://leetcode.com/problems/car-fleet/


class SimplerSolution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        pair = [(p, s) for p, s in zip(position, speed)]
        pair.sort()

        stack = []
        for i in range(len(pair) - 1, -1, -1):
            position, speed = pair[i][0], pair[i][1]
            timeToReach = (target - position) / speed

            if len(stack) == 0 or timeToReach > stack[-1]:
                stack.append(timeToReach)

        return len(stack)


class Solution:
    def carFleet(self, target, position, speed):
        answer = 0
        timeToDest = []

        pair = [(p, s) for p, s in zip(position, speed)]
        pair.sort(reverse=True)

        stack = []
        for position, speed in pair:
            time = (target - position) / speed
            if len(stack) == 0:
                stack.append(time)
                continue

            top = stack[-1]
            if time > top:
                stack.append(time)

        return len(stack)


target = 12
position = [10, 8, 0, 5, 3]
speed = [2, 4, 1, 1, 3]

s = Solution()
print("-> ", s.carFleet(target, position, speed))
