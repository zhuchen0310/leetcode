import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	for len(stones) > 1 {
		a := stones[len(stones)-1]
		b := stones[len(stones)-2]
		stones = stones[:len(stones)-2]
		if a > b {
			stones = append(stones, a-b)
			sort.Ints(stones)
		}
	}
	if len(stones) > 0 {
		return stones[0]
	} else {
		return 0
	}
}