# https://leetcode.com/problems/non-overlapping-intervals/
#
# O(n lon n) time
# O(1) memory
class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda i: i[0])
        result = 0
        prevEnd = intervals[0][1]

        for i in range(1, len(intervals)):
            if prevEnd <= intervals[i][0]:
                prevEnd = intervals[i][1]
                continue
            else:
                prevEnd = min(prevEnd, intervals[i][1])
                result += 1

        return result
