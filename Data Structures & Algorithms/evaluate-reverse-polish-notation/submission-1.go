func Calculate(number2 string, number1 string, operator string) string {
	num1, _ := strconv.Atoi(number1)
	num2, _ := strconv.Atoi(number2)
	if operator == "+" {
		return strconv.Itoa(num1 + num2)
	} else if operator == "*" {
		return strconv.Itoa(num1 * num2)
	} else if operator == "/" {
		return strconv.Itoa(num1 / num2)
	} else {
		return strconv.Itoa(num1 - num2)
	}
}
func evalRPN(tokens []string) int {
	stack := make([]string, 0)
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
		stack = append(stack, token)
	}
	resultString := stack[len(stack)-1]
	result, _ := strconv.Atoi(resultString)
	return result
}
