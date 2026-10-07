class Solution:
    def isValid(self, i, k, board: list[list[str]]) -> bool:
        hashmap={}
        for r in range(i, i+3):
            for c in range(k, k+3):
                if board[r][c] == '.':
                    continue
                if board[r][c] in hashmap:
                    return False
                hashmap[board[r][c]]=True
        return True
    def isValidSudoku(self, board: list[list[str]]) -> bool:
        for i in range(0, len(board)):
            hashmap={}
            for k in range(0, len(board)):
                if board[i][k] == '.':
                    continue
                if board[i][k] in hashmap:
                    return False
                hashmap[board[i][k]]=True
        for i in range(0, len(board)):
            hashmap={}
            for k in range(0, len(board)):
                if board[k][i] == '.':
                    continue
                if board[k][i] in hashmap:
                    return False
                hashmap[board[k][i]]=True

        for i in range(0, len(board), 3):
            for k in range(0, len(board[i]), 3):
                if not self.isValid(i, k, board):
                    return False
        return True