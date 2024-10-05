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


# Time: O(n) avg case; O(n^2) worst case (when the pivot is always the greatest)
# Memory: O(n)
class QuickSelectSolution:
    def partition(self, nums, left, right):
        pivot, divisionIndex = nums[right], left
        for i in range(left, right):
            if nums[i] < pivot:
                nums[i], nums[divisionIndex] = nums[divisionIndex], nums[i]
                divisionIndex += 1

        nums[right], nums[divisionIndex] = nums[divisionIndex], nums[right]
        return divisionIndex

    def select(self, nums, k, left, right):
        pivotIndex = self.partition(nums, left, right)

        if k == pivotIndex + 1:  # found
            return nums[pivotIndex]
        elif k < pivotIndex + 1:  # search in the smallest array
            return self.select(nums, k, left, pivotIndex - 1)
        else:  # search in the biggest array
            return self.select(nums, k, pivotIndex + 1, right)

    def findKthLargest(self, nums: List[int], k: int) -> int:
        return self.select(nums, len(nums) - k + 1, 0, len(nums) - 1)
