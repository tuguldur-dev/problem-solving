func checkInclusion(s1 string, s2 string) bool {
	l := len(s1)
	s1a := []byte(s1)
	sort.Slice(s1a, func(a, b int) bool {
		return s1a[a] < s1a[b]
	})
	for i := 0; i <= len(s2)-l; i++ {
		bb := []byte(s2[i:i+l])
		sort.Slice(bb, func(a,b int) bool {
			return bb[a] < bb[b]
		})
		if string(s1a) == string(bb) {
			return true
		}
	}
	return false
}
