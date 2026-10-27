func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	longest := 0
	maxFreq := 0
	start := 0
	for end := 0; end < len(s); end++ {
		freq[s[end]]++
		maxFreq=max(maxFreq, freq[s[end]])
		for (end-start+1)-maxFreq > k {
			freq[s[start]]--
			start++
		}
		longest=max(longest,end-start+1)
	}
	return longest
}
