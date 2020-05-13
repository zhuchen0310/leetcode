class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        stones_heap = [-i for i in stones]
        heapq.heapify(stones_heap)
        while len(stones_heap) > 1:
            a = heapq.heappop(stones_heap)
            b = heapq.heappop(stones_heap)
            if a < b:  # 如果 第一块石头终于第二块
                heapq.heappush(stones_heap, a-b)
        if stones_heap:
            return -stones_heap[0]
        return 0
