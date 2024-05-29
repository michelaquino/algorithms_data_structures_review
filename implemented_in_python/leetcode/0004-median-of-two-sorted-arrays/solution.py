# https://leetcode.com/problems/median-of-two-sorted-arrays/


class Solution(object):

    # (m + n log(m + n ))
    def findMedianSortedArraysNaive(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """

        merged = nums1 + nums2
        merged.sort()

        lastIndex = len(merged) - 1
        if len(merged) % 2 == 0:
            first = float(merged[lastIndex // 2])
            second = float(merged[(lastIndex // 2) + 1])
            return (first + second) / 2
        else:
            return merged[float(lastIndex) // 2]

    def findMedianSortedArrays(self, nums1, nums2):
        if len(nums1) <= len(nums2):
            return self.find(nums1, nums2)
        else:
            return self.find(nums2, nums1)

    def find(self, smallest, biggest):
        total = len(smallest) + len(biggest)
        half = total // 2

        left, right = 0, len(smallest) - 1
        while True:
            middle = left + ((right - left) // 2)
            lastIndexBigger = half - middle - 2

            smallestLeft = smallest[middle] if middle >= 0 else float("-infinity")
            smallestRight = (
                smallest[middle + 1]
                if middle + 1 < len(smallest)
                else float("infinity")
            )

            biggestLeft = (
                biggest[lastIndexBigger] if lastIndexBigger >= 0 else float("-infinity")
            )

            biggestRight = (
                biggest[lastIndexBigger + 1]
                if lastIndexBigger + 1 < len(biggest)
                else float("infinity")
            )

            if smallestLeft > biggestRight:
                right = middle - 1
            elif biggestLeft > smallestRight:
                left = middle + 1
            else:
                # odd
                if total % 2:
                    return min(smallestRight, biggestRight)
                # even
                return (
                    max(smallestLeft, biggestLeft) + min(smallestRight, biggestRight)
                ) / 2


s = Solution()
nums1 = [1, 3]
nums2 = [2]
print("2 -->", s.findMedianSortedArrays(nums1, nums2))


nums1 = [1, 2]
nums2 = [3, 4]
print("2.5 --> ", s.findMedianSortedArrays(nums1, nums2))

nums1 = [1, 2, 3, 4]
nums2 = [5, 6, 7]
print("4 -->", s.findMedianSortedArrays(nums1, nums2))

nums1 = [4, 5, 6, 7]
nums2 = [1, 2, 3]
print("4 -->", s.findMedianSortedArrays(nums1, nums2))

nums1 = [1, 2, 6, 7]
nums2 = [3, 4, 5]
print("4 -->", s.findMedianSortedArrays(nums1, nums2))

nums1 = [1, 2, 3, 4, 5]
nums2 = [6, 7, 8]
print("4.5 -->", s.findMedianSortedArrays(nums1, nums2))

nums1 = [4, 5, 6, 7, 8]
nums2 = [1, 2, 3]
print("4.5 -->", s.findMedianSortedArrays(nums1, nums2))

nums1 = [1, 2, 6, 7, 8]
nums2 = [3, 4, 5]
print("4.5 -->", s.findMedianSortedArrays(nums1, nums2))
