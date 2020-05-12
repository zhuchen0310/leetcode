func buildArray(target []int, n int) []string {
	ret := []string{}
	pop := "Pop"
	push := "Push"
	if len(target) == 0 {
		for i := 1; i <= n; i++ {
			ret = append(ret, push, pop)
		}
	}
	stack := []int{}
	start := 0
	for j := 1; j <= n; j++ {
		if len(stack) == len(target) {
			break
		}
		ret = append(ret, push)
		if target[start] != j {
			ret = append(ret, pop)
		} else {
			start++
			stack = append(stack, j)
		}
	}
	return ret
}