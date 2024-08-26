# https://leetcode.com/problems/k-closest-points-to-origin/


# O(n) memory
# O(k log n)
class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        heap = []

        # O(n)
        for point in points:
            x, y = point[0], point[1]
            distance = math.sqrt(math.pow(x, 2) + math.pow(y, 2))
            heap.append((distance, point))

        heapq.heapify(heap)  # O(n)

        result = []

        # O(k)
        for i in range(0, k):
            closest = heapq.heappop(heap)
            result.append(closest[1])  # O(log n)

        return result
