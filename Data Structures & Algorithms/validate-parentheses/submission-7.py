class Solution:
    def isValid(self, s: str) -> bool:
        stack = []

        for char in s:
            if char == '[' or char == '(' or char == '{':
                stack.append(char)
                continue
            if not stack:
                return False
            if char == ']' and stack[-1] == '[':
                stack.pop()
                continue
            elif char == ')' and stack[-1] == '(':
                stack.pop()
                continue
            elif char == '}' and stack[-1] == '{':
                stack.pop()
                continue
            else:
                return False


        return len(stack) == 0

            