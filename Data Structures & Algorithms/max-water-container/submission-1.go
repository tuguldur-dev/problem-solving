func maxArea(heights []int) int {
	left := 0
	right := len(heights)-1
	maxArea := 0
	for left < right {
		maxArea = max(maxArea, min(heights[left], heights[right]) * (right-left))
		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
