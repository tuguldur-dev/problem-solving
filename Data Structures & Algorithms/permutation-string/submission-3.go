func checkInclusion(s1 string, s2 string) bool {
	l := len(s1)
	if l > len(s2) {
		return false
	}
	var need [26]int
	var window [26]int

	for i := 0; i < l; i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	if need == window {
		return true
	}
	for i := l; i < len(s2); i++ {
		window[s2[i]-'a']++
		window[s2[i-l]-'a']--

		if window == need {
			return true
		}
	}
	return false

}
