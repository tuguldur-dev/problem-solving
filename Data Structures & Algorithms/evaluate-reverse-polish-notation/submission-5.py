class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        operators ={'*', '-','+','/'}
        stack = []
        for token in tokens:
            if token not in operators:
                stack.append(int(token))
                continue
            b, a = stack.pop(), stack.pop()
            result = 0
            if token == '*':
                result = a * b
            elif token == '/':
                result = a / b
            elif token == '+':
                result = a + b
            else:
                result = a - b
            stack.append(int(result))
        return stack[0]
        