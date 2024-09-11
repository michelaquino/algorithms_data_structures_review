# https://leetcode.com/problems/gas-station/


class Solution:
    # Greedy
    # O(n) time
    # O(1) memory
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        if sum(gas) < sum(cost):
            return -1

        result = 0
        total = 0
        for i in range(len(gas)):
            total += gas[i] - cost[i]
            if total < 0:
                total = 0
                result = i + 1

        return result

    # O(n^2) time
    # O(1) memory
    def canCompleteCircuit_brute_force(self, gas: List[int], cost: List[int]) -> int:
        result = -1
        for i in range(len(gas)):
            if self.wasCircuitCompleted(i, gas, cost):
                return i

        return result

    def wasCircuitCompleted(self, startStation, gas, cost):
        tank = gas[startStation]

        station = startStation % len(gas)
        for j in range(len(gas)):
            tank = tank - cost[station]
            if tank < 0:
                return False

            station = (station + 1) % len(gas)
            tank += gas[station]

        return True
