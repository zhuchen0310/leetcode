class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        if not nums:
            return -1 
        m = {}
        for n in nums:
            if n in m:
                return n 
            else:
                m[n] = 0
        return -1 