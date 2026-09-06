func countBits(n int) []int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = dp[i/2] + (i%2)
	}
	return dp
}
