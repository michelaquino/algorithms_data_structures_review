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


# Time: O(n * m) where m is the iddle time
# Memory: O(n)
class NeetcodeSolution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        count = Counter(tasks)
        maxHeap = [-cnt for cnt in count.values()]
        heapq.heapify(maxHeap)

        time = 0
        q = deque()  # pairs of [-cnt, idleTime]
        while maxHeap or q:
            time += 1

            if not maxHeap:
                time = q[0][1]
            else:
                cnt = 1 + heapq.heappop(maxHeap)
                if cnt:
                    q.append([cnt, time + n])
            if q and q[0][1] == time:
                heapq.heappush(maxHeap, q.popleft()[0])
        return time
