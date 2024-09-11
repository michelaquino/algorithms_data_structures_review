class Solution:
    # O(n) time
    # O(1) memory
    def canJump(self, nums: List[int]) -> bool:
        goal = len(nums) - 1
        for i in range(len(nums) - 1, -1, -1):
            if i + nums[i] >= goal:
                goal = i

        return goal == 0

    # O(n^2) time, because the constraints: nums[i] can be greater than nums.length
    # O(n) memory
    def canJumpDP(self, nums: List[int]) -> bool:
        dp = [False] * len(nums)
        dp[len(nums) - 1] = True

        for i in range(len(nums) - 1, -1, -1):
            reachEnd = False
            for j in range(0, nums[i] + 1):
                nextPosition = i + j
                if nextPosition >= len(nums) - 1 or dp[nextPosition]:
                    reachEnd = True
                    break

            dp[i] = reachEnd
        return dp[0]
