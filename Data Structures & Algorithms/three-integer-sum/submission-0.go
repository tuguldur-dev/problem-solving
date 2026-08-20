func threeSum(nums []int) [][]int {
	sort.Ints(nums)	

	// [-1,0,1,2,-1,-4]
	// [-4,-1,-1,0,1,2]
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i+1
		right := len(nums)-1
		for left < right {
			if -nums[i] < nums[left]+nums[right] {
				right--
			} else if -nums[i] > nums[left]+nums[right] {
				left++
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++	
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}
