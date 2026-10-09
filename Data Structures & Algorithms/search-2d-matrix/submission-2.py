class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        top, bot = 0, len(matrix)-1

        while top <= bot:
            row_mid = int((top+bot)/2)
            if matrix[row_mid][-1] < target:
                top = row_mid+1
            elif matrix[row_mid][0] > target:
                bot = row_mid-1
            else: 
                left, right = 0, len(matrix[row_mid])-1
                while left <= right:
                    col_mid = int((right + left) / 2)
                    if matrix[row_mid][col_mid] < target:
                        left = col_mid+1
                    elif matrix[row_mid][col_mid] > target:
                        right=col_mid-1
                    else:
                        return True
                return False
        return False
                


