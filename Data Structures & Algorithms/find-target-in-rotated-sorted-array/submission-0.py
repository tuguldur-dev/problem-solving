class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l,r=0, len(nums)-1
        while l < r:
            m=(l+r)//2
            if nums[m] > nums[r]:
                l=m+1
            else:
                r=m
        rotation=l

        l,r=0, len(nums)-1
        while l <= r:
            m=(l+r)//2
            actual=(m+rotation) % len(nums)
            if nums[actual] > target:
                r=m-1
            elif nums[actual] < target:
                l=m+1
            else:
                return actual
        return -1

