func lengthOfLongestSubstring(s string) int {
	check := make(map[byte]int) 
	start := 0
	longest := 0
	for end := 0; end < len(s); end++ {
		check[s[end]]++

		for start < end && check[s[end]] > 1 {
			check[s[start]]--
			start++
		}

		longest = max(longest, end-start+1)
	}
	return longest
}
