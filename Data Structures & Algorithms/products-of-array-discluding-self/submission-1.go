func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	prefix := 1
	for i := 0; i < len(nums); i++ {
		output[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := len(nums)-1; i >= 0; i-- {
		output[i] *= suffix
		suffix *= nums[i]
	}

	return output
}
