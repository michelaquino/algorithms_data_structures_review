# https://leetcode.com/problems/task-scheduler/


# Priority Queue solution
# Time: O(n log n)
# Memory: O(n)
class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        taskPriority = {}
        queue = []

        for task in tasks:  # O(n)
            priority = 1
            if task in taskPriority:  # O(1)
                priority = taskPriority[task] + n + 1

            taskPriority[task] = priority  # O(1)
            heapq.heappush(queue, priority)  # O(log n)

        answer = 0
        interval = 1
        while queue:  # O(n)
            if queue[0] <= interval:
                heapq.heappop(queue)  # O(log n)

            answer += 1
            interval += 1
        return answer
