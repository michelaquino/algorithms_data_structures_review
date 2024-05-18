# https://leetcode.com/problems/best-time-to-buy-and-sell-stock

class Solution:
    def maxProfit(self, prices):
        maxValue = 0
        buyDay, sellDay = 0, 1

        while sellDay < len(prices):
            maxValue = max(maxValue, prices[sellDay] - prices[buyDay])
            
            if prices[buyDay] < prices[sellDay]:
                sellDay += 1
            else:
                buyDay = sellDay
                sellDay += 1

        return max(maxValue, 0)
    
s = Solution()
# print(s.maxProfit([]))
# print(s.maxProfit([1, 1, 1, 1, 1]))
# print(s.maxProfit([7,1,5,3,6,4]))
# print(s.maxProfit([7,6,4,3,1]))
# print(s.maxProfit([1, 2, 3, 4, 5, 6]))
# print(s.maxProfit([6, 5, 4, 3, 2, 1]))
# print(s.maxProfit([6, 5, 4, 3, 2, 7]))
# print(s.maxProfit([6, 6, 6, 6, 6, 7]))
# print(s.maxProfit([2,1,2,1,0,1,2]))
print(s.maxProfit([1,2,4,2,5,7,2,4,9,0,9]))