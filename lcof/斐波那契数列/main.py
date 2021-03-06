class Solution:
    @functools.lru_cache()
    def fib(self, n: int) -> int:
        if n < 2:
            return n
        return (self.fib(n-1) + self.fib(n-2)) % 1000000007
