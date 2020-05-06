func isValid(s string) bool {
	if s == "" {
		return true
	}
	sMap := map[byte]byte{')': '(', '}': '{', ']': '['}
	stack := []byte{}
	if _, ok := sMap[s[0]]; ok {
		return false
	}
	// 需要将字符串都专程byte 数组
	newS := []byte(s)
	for _, sub := range newS {
		vs, ok := sMap[sub]
		if len(stack) > 0 && ok && stack[len(stack)-1] == vs {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, sub)
		}
	}
	return len(stack) == 0
}