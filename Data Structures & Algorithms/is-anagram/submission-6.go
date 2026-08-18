func isAnagram(s string, t string) bool {
	hashmap := make(map[rune]int)

	for _, c1 := range s {
		hashmap[c1]++
	}

	for _, c2 := range t {
		_, ok := hashmap[c2]
		if !ok {
			return false
		}
		hashmap[c2]--
		if hashmap[c2] <= 0 {
			delete(hashmap, c2)
		}

	}
	return len(hashmap) == 0
}
