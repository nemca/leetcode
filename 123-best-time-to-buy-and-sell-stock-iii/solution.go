package best_time_to_buy_and_sell_stock_iii

import "math"

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	buy1, sell1 := math.MaxInt32, 0
	buy2, sell2 := math.MaxInt32, 0

	for _, price := range prices {
		buy1 = min(buy1, price)
		sell1 = max(sell1, price-buy1)

		buy2 = min(buy2, price-sell1)
		sell2 = max(sell2, price-buy2)
	}
	return sell2
}
