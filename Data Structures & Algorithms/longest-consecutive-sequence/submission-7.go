func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	longest := 1
	current := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]{
			continue
		} 
		if nums[i-1]+1 == nums[i] {
			current++
		} else {
			current=1
		}
		longest = max(current, longest)
	}
	return longest
}
