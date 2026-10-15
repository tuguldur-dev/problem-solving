class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        start = 0
        end = 0
        mprofit = 0
        while end < len(prices):
            profit = prices[end]-prices[start]
            mprofit = max(mprofit, profit)
            if prices[end] < prices[start]:
                start = end
            end += 1
        return mprofit
        
        