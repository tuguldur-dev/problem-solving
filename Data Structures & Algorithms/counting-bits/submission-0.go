func countBits(n int) []int {
	result := make([]int, 0)

	for i := 0; i <= n; i++ {
		k := i
		counter := 0
		for k > 0 {
			if k % 2 > 0 {
				counter++
			}
			k = k / 2
		}
		result = append(result, counter)
	}
	return result
}
