class Solution:
    def isValid(self, s: str) -> bool:
        operators = {']':'[', ')':'(', '}':'{'}
        stack = []

        for c in s:
            if c not in operators:
                stack.append(c)
                continue
            if not stack:
                return False
            if stack[-1] == operators[c]:
                stack.pop()
            else:
                return False
        return len(stack) == 0
        