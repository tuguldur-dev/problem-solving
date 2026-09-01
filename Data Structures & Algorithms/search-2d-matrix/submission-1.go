func searchMatrix(matrix [][]int, target int) bool {
	left := 0
	right := len(matrix)-1
	colMid := (left+right)/2
	for left <= right {
		colMid = (left+right)/2
		fmt.Println(colMid)

		if matrix[colMid][0] > target {
			right = colMid-1
		} else if matrix[colMid][len(matrix[colMid])-1] < target {
			left = colMid+1	
		} else {
			break
		}
	}

	left = 0
	right = len(matrix[colMid])-1

	for left <= right {
		mid := (left+right)/2
		if matrix[colMid][mid] > target {
			right = mid - 1
		} else if matrix[colMid][mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
