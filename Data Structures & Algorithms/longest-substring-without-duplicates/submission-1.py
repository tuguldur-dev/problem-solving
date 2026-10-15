class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        count={}
        start = 0
        longest = 0
        for end in range(len(s)):
            count[s[end]] = count.get(s[end], 0) + 1
            while count.get(s[end], 0) > 1:
                count[s[start]] -= 1
                start += 1
            longest = max(longest, end-start+1)
        return longest
