# https://leetcode.com/problems/merge-intervals/
#
# O(n log n) time
# O(n) memory
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda pair: pair[0])
        result = []

        interval = intervals[0]
        for i in range(1, len(intervals)):

            # left
            if interval[1] < intervals[i][0]:
                result.append(interval)
                interval = intervals[i]
            # right
            elif interval[0] > intervals[i][1]:
                result.append(intervals[i])

            # merge
            else:
                interval = [
                    min(interval[0], intervals[i][0]),
                    max(interval[1], intervals[i][1]),
                ]

        result.append(interval)
        return result
