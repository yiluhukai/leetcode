package  main

import "fmt"

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。 
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。 
//
// 
//
// 示例 1: 
//
// 输入: [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
// 
//
// 示例 2: 
//
// 输入: [1,2,3,4,5]
//输出: 4
//解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
// 
//
// 示例 3: 
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。 
//
// 
//
// 提示： 
//
// 
// 1 <= prices.length <= 3 * 10 ^ 4 
// 0 <= prices[i] <= 10 ^ 4 
// 
// Related Topics 贪心算法 数组 
// 👍 1069 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
	使用贪心算法，获取当前的最大值，也就是每一步
	执行耗时:4 ms,击败了94.37% 的Go用户
    内存消耗:3 MB,击败了100.00% 的Go用户

	开始增长的时候买入，开始减少的时候卖出，买入前需要检查是否买入过，卖出前需要检查是否卖出过
 */
func maxProfit(prices []int) int {
	length:= len(prices)
	buyIndex :=-1
	// 利润
	profit := 0
	for index,price :=range prices{
		// 开始增长的时候买入
		if index+1 <length && price < prices[index+1]{
			// 还没有买入
			if buyIndex ==-1{
				buyIndex = index
			}
		}else if (index+1 <length && price > prices[index+1]) || index+1 == length{
			//开始递减或者到达末尾的时候卖出且已经买过
			if buyIndex !=-1{
				profit = profit + price - prices[buyIndex]
				buyIndex = -1
			}
		}
	}
	return profit
}
//leetcode submit region end(Prohibit modification and deletion)

func main(){
	res:= maxProfit([]int{1,2,3,4,5})
	fmt.Printf("%v\n",res)
}
