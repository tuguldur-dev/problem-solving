class Solution:
    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        hashmap={}
        for num in nums:
            hashmap[num]=hashmap.get(num, 0)+1
        unique=list(hashmap.keys())
        unique.sort(key=lambda k: hashmap[k], reverse=True)
        return unique[:k]
            
        