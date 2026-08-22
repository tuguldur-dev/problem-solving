func maxProfit(prices []int) int {
	buy := 0
	maxProfit := 0
	for sell := 1; sell < len(prices); sell++ {
		profit := prices[sell] - prices[buy]
		maxProfit = max(maxProfit, profit)

		if prices[buy] > prices[sell] {
			buy=sell
		}
	}
	return maxProfit
}
