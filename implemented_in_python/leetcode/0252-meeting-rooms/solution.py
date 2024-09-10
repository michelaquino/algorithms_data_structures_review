"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""


# https://neetcode.io/problems/meeting-schedule
# https://leetcode.com/problems/meeting-rooms/ (premium)
#
# O(n log n) time
# O(1) memory
class Solution:
    def canAttendMeetings(self, intervals: List[Interval]) -> bool:
        if len(intervals) == 0:
            return True

        intervals.sort(key=lambda i: i.start)
        interval = intervals[0]
        for i in range(1, len(intervals)):
            if intervals[i].start < interval.end:
                return False

            interval = intervals[i]

        return True
