class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        hashmap={}
        for index, num in enumerate(nums):
            c = target-num
            if c in hashmap:
                return [hashmap[c], index]
            hashmap[num]=index
        return []
        