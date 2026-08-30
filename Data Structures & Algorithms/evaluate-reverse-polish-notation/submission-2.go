func Calculate(number2 int, number1 int, operator string) int {
	if operator == "+" {
		return number1 + number2
	} else if operator == "*" {
		return number1 * number2
	} else if operator == "/" {
		return number1 / number2
	} else {
		return number1 - number2
	}
}
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			number1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			number2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result := Calculate(number1, number2, token)
			stack = append(stack, result)
			continue
		}
		num,_ := strconv.Atoi(token)
		stack = append(stack, num)
	}
	return stack[0]
}
