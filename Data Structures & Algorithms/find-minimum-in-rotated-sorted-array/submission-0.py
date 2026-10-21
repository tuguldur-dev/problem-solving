class Solution:
    def findMin(self, nums: List[int]) -> int:
        minNum = float('inf')
        for num in nums:
            minNum = min(num, minNum)
        return minNum