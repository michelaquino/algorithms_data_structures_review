class Solution(object):
    # two pointers solution
    # O(n) time complexity
    # No extra space complexity
    def twoSum(self, numbers, target):
        """
        :type numbers: List[int]
        :type target: int
        :rtype: List[int]
        """

        answer = []
        left, right = 0, len(numbers) - 1

        while left < right:
            sumResult = numbers[left] + numbers[right]
            if sumResult == target:
                return [left + 1, right + 1]

            if sumResult > target:
                right -= 1
            else:
                left += 1

        return []

    # hash table solution
    # O(n) time complexity
    # Extra space complexity
    def twoSumHashSolution(self, numbers, target):
        """
        :type numbers: List[int]
        :type target: int
        :rtype: List[int]
        """

        numberIndexHash = {}
        for index, number in enumerate(numbers):
            numberIndexHash[number] = index

        # print(numberIndexHash)
        answer = []

        for index, number in enumerate(numbers):
            delta = target - number

            if delta in numberIndexHash and numberIndexHash[delta] > index:
                return [index + 1, numberIndexHash[delta] + 1]

        return []


s = Solution()

numbers = [2, 7, 11, 15]
target = 9
print(s.twoSum(numbers, target))

numbers = [2, 3, 4]
target = 6
print(s.twoSum(numbers, target))

numbers = [-1, 0]
target = -1
print(s.twoSum(numbers, target))
