# https://en.wikipedia.org/wiki/Quickselect
#
# Example to use: Find the kth smallest/greatest element in the array
#
# Time: O(n) in the average case; O(n^2) worst case (when the pivot is always the greatest)
# Memory: O(n)
class QuickSelect:
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

    def selectKhSmallest(self, nums, k):
        return self.select(nums, k, 0, len(nums) - 1)

    def selectKhGreatest(self, nums, k):
        return self.select(nums, len(nums) - k + 1, 0, len(nums) - 1)


q = QuickSelect()

nums = [3, 2, 1, 5, 6, 4]
print("1 smallest: ", q.selectKhSmallest(nums.copy(), 1))
print("2 smallest: ", q.selectKhSmallest(nums.copy(), 2))
print("3 smallest: ", q.selectKhSmallest(nums.copy(), 3))
print("4 smallest: ", q.selectKhSmallest(nums.copy(), 4))
print("5 smallest: ", q.selectKhSmallest(nums.copy(), 5))
print("6 smallest: ", q.selectKhSmallest(nums.copy(), 6))
print()
print("1 greatest: ", q.selectKhGreatest(nums.copy(), 1))
print("2 greatest: ", q.selectKhGreatest(nums.copy(), 2))
print("3 greatest: ", q.selectKhGreatest(nums.copy(), 3))
print("4 greatest: ", q.selectKhGreatest(nums.copy(), 4))
print("5 greatest: ", q.selectKhGreatest(nums.copy(), 5))
print("6 greatest: ", q.selectKhGreatest(nums.copy(), 6))
