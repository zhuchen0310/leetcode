func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// 小于第一个则肯定不存在
	if target < matrix[0][0] {
		return false
	}
	x := len(matrix[0]) - 1
	y := 0
	// 开始从第一行最右侧遍历。 大于目标值则向下 y++， 否则向左x--
	for x >= 0 && y < len(matrix) {
		if target == matrix[y][x] {
			return true
		} else if target > matrix[y][x] {
			y++
		} else {
			x--
		}
	}
	return false
}