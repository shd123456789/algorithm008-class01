// 时间复杂度为O(n) 空间复杂度为O(1) 2遍
func maxProfit(prices []int) int {
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i - 1] {
            maxProfit += prices[i] - prices[i - 1]
        }
    }
    return maxProfit
}