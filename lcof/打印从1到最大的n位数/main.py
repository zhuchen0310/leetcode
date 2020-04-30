class Solution:
    def printNumbers(self, n: int) -> List[int]:
        s = 1
        while n > 0:
            s *= 10
            n -= 1
        return [x for x in range(1, s)]