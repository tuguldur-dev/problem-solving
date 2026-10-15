class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        start = 0
        profit = 0
        for end in range(1, len(prices)):
            if prices[end] < prices[start]:
                start = end
            else:
                profit=max(profit, prices[end]-prices[start])
        return profit
        
        