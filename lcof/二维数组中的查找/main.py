class Solution:
    def findNumberIn2DArray(self, matrix: List[List[int]], target: int) -> bool:
        if not matrix or not matrix[0]:
            return False
        x = len(matrix[0])-1
        y = 0 
        while x >=0 and y < len(matrix):
            print(matrix[y][x])
            if target == matrix[y][x]:
                return True
            # 如果target 大于最右侧 则向下 y+=1
            elif target > matrix[y][x]:
                y += 1
            # 否则target 小于当前值，则向左 x-=1
            else:
                x -= 1
        return False 