class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        operation=set(['+','-','/','*'])
        for token in tokens:
            if token not in operation:
                stack.append(int(token))
            else:
                result = 0
                b=stack.pop()
                a=stack.pop()
                if token == '*': 
                    result = a * b
                elif token == '/':
                    result = a / b
                elif token == '-':
                    result = a - b
                else:
                    result = a + b
                stack.append(int(result))
        return stack[0]

        