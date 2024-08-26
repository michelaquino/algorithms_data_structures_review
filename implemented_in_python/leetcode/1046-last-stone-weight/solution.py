# https://leetcode.com/problems/last-stone-weight/
# O(n log n) time
# O(n) memory


class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:

        heap = []
        heapq.heapify(heap)
        for s in stones:  # n
            heapq.heappush(heap, s * -1)  # log n

        while len(heap) > 1:
            y = heapq.heappop(heap)  # log n
            x = heapq.heappop(heap)  # log n

            if x != y:
                y -= x
                heapq.heappush(heap, y)  # log n

        return heap[0] * -1 if len(heap) > 0 else 0
