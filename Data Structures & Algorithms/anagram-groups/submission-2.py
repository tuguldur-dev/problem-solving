class Solution:
    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        group={}
        for st in strs:
            key=tuple(sorted(st))
            group.setdefault(key, []).append(st)

        return list(group.values())
