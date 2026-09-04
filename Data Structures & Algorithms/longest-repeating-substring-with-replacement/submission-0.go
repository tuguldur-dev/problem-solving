func characterReplacement(s string, k int) int {
	counting := make(map[byte]int)
	start := 0
	maxFreq := 0
	longestRepeat := 0

	for end := 0; end < len(s); end++ {
		counting[s[end]]++
		maxFreq = max(maxFreq, counting[s[end]])

		for (end-start+1)-maxFreq > k {
			counting[s[start]]--
			start++
		}
		longestRepeat = max(longestRepeat, end-start+1)
	}
	return longestRepeat
}
