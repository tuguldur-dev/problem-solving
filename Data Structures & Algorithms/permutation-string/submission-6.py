class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        l=len(s1)
        if l > len(s2):
            return False
        need=26*[0]
        window=26*[0]
        for i in range(l):
            need[ord(s1[i])-ord('a')]+=1
            window[ord(s2[i])-ord('a')]+=1
        if window == need: 
            return True

        for i in range(l, len(s2)):
            window[ord(s2[i])-ord('a')]+=1
            window[ord(s2[i-l])-ord('a')]-=1
            if window == need:
                return True
        return False
        