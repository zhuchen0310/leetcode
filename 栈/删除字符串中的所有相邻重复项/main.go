func removeDuplicates(S string) string {
	stack := []byte{}
	for _, s := range []byte(S) {
		if len(stack) == 0 {
			stack = append(stack, s)
		} else {
			if s == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s)
			}
		}
	}
	return string(stack)
}