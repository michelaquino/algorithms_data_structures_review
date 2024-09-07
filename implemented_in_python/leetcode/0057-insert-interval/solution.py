# https://leetcode.com/problems/insert-interval/
#
# O(n) time
# O(n) memory
class Solution:
    def insert(
        self, intervals: List[List[int]], newInterval: List[int]
    ) -> List[List[int]]:
        result = []

        for i, interval in enumerate(intervals):
            # left
            if newInterval[1] < interval[0]:
                result.append(newInterval)
                return result + intervals[i:]

            # right
            elif newInterval[0] > interval[1]:
                result.append(interval)

            # merge
            else:
                newInterval = [
                    min(newInterval[0], interval[0]),
                    max(newInterval[1], interval[1]),
                ]

        result.append(newInterval)
        return result
