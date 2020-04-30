func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	end := len(numbers) - 1
	// 从右侧开始遍历， 如果是反转的列表那肯定小于左侧数字
	// 注意边界
	for ; end > 0; end-- {
		if numbers[end-1] > numbers[end] {
			return numbers[end]
		}
	}
	return numbers[end]
}