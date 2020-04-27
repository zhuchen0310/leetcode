package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	max := 0
	for _, p := range prices[1:] {
		if p < buy {
			buy = p
		}
		if p-buy > max {
			max = p - buy
		}
	}

	return max
}
