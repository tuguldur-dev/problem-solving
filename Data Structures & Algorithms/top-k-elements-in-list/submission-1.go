func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)
	for _, num := range nums {
		hashmap[num]++
	}
    unique := make([]int, 0, len(hashmap))

	for num := range hashmap {
		unique = append(unique, num)
	}

	sort.Slice(unique, func(i, j int) bool {
        return hashmap[unique[i]] > hashmap[unique[j]]
    })

	
	return unique[:k]
}
