//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

//解题思路
//把买卖股票拆分成小块，如果跌就不操作，涨就模拟单天买入卖出，一次循环解决

func maxProfit(prices []int) int {
	maxPro := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			maxPro += prices[i] - prices[i-1]
		}
	}
	return maxPro
}