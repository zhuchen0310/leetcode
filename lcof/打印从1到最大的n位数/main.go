func printNumbers(n int) []int {
	ret := []int{}
	sum := 1
	for n > 0 {
		sum = sum * 10
		n--
	}
	for i := 1; i < sum; i++ {
		ret = append(ret, i)
	}
	return ret
}