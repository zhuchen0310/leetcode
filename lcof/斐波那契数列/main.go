var numMap = make(map[int]int)

func fib(n int) int {
	if n < 2 {
		return n
	}
	if v, ok := numMap[n]; ok {
		return v
	} else {
		num := (fib(n-1) + fib(n-2)) % 1000000007
		numMap[n] = num
		return num
	}
}