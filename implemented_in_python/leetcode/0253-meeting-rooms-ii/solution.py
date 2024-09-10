"""
Definition of Interval:
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""

# O(n log n) time
# O(n) memory


class Solution:
    def minMeetingRooms(self, intervals: List[Interval]) -> int:
        start, end = [], []

        for i in intervals:
            start.append(i.start)
            end.append(i.end)

        start.sort()
        end.sort()

        maxCount, count = 0, 0
        s, e = 0, 0
        while s < len(start):
            if start[s] < end[e]:
                s += 1
                count += 1
            else:
                e += 1
                count -= 1

            maxCount = max(maxCount, count)

        return maxCount
