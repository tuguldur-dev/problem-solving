func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			stack = append(stack, c)
			continue
		}  
		if len(stack) > 0 {
			last := len(stack)-1
			if stack[last] == '[' && c == ']' {
				stack = stack[:last]
			} else if stack[last] == '{' && c == '}' {
				stack = stack[:last]
			} else if stack[last] == '(' && c == ')' {
				stack = stack[:last]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}
