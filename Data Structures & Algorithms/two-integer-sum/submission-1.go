func twoSum(nums []int, target int) []int {
    hashmap := make(map[int]int)

	for index, num := range nums {

		if _index, ok := hashmap[num]; ok && _index != index {
			return []int{_index, index}
		}
		hashmap[target-num]=index

	}
	return []int{}
}
