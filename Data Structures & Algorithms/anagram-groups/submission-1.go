func groupAnagrams(strs []string) [][]string {
	hashmap := make(map[string][]string)

	for _, str := range strs {
		keyBytes := []byte(str)
		sort.Slice(keyBytes, func(i, j int) bool{
			return keyBytes[i] < keyBytes[j]
		})
		key := string(keyBytes)
		hashmap[key] = append(hashmap[key], str)
	}
	
	result := make([][]string, 0, len(hashmap))

	for _, grp := range hashmap {
		result = append(result, grp)
	}
	return result
}


