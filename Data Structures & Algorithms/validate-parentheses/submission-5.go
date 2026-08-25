func isValid(s string) bool {
	stack := make([]rune, 0)
	pairs := map[rune]rune {
		'}': '{',
		']': '[',
		')': '(',
	}


	for _, c := range s {
		if opening, ok := pairs[c];  ok {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
