func removeOuterParentheses(S string) string {
	if len(S) == 0 {
		return ""
	}
	ret := ""
	stack := []byte{S[0]}
	for start := 1; start < len(S); start++ {
		s := S[start]
		if s == '(' {
			stack = append(stack, s)
			if len(stack) > 1 {
				ret += string(s)
			}
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				ret += string(s)
			}
		}
	}
	return ret
}