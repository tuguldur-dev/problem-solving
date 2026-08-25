func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c)
			continue
		}  
		if len(stack) > 0 {
			if stack[len(stack)-1] == '[' && c == ']' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] == '{' && c == '}' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] == '(' && c == ')' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}
