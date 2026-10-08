class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        operators = {']': '[', '}': '{', ')': '('}
        for char in s:
            if char in operators.values():
                stack.append(char)
                continue
            if len(stack) == 0:
                return False
            if operators[char] == stack[-1]:
                stack.pop()
            else:
                return False

        return len(stack) == 0

            