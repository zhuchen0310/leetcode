func backspaceCompare(S string, T string) bool {
	return parseStr(S) == parseStr(T)
}

func parseStr(Str string) string {
	stack := []byte{}
	for _, s := range []byte(Str) {
		if s == '#' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if s != '#' {
			stack = append(stack, s)
		}
	}
	ret := string(stack)
	return ret
}