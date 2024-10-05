# https://leetcode.com/problems/kth-largest-element-in-an-array/


# Time: O(n + k log n)
# Memory: O(1)
class MaxHeapSolution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        for i in range(len(nums)):  # n
            nums[i] *= -1

        heapq.heapify(nums)  # n
        largest = 0
        while k > 0:  # k
            largest = heapq.heappop(nums)  # log n
            k -= 1

        return largest * -1
