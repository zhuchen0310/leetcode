class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if not prices:
            return 0
        buy = prices[0] # 首次买入价格
        m = 0   # 当前最大收益
        for p in prices[1:]:
            # 每次获取最小买入价格
            buy = min([buy, p])
            # 每次获取最大收益
            m = max([p-buy, m])
        return m