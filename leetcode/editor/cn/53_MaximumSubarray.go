//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。 
//
// 示例: 
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 
//
// 进阶: 
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。 
// Related Topics 数组 分治算法 动态规划 
// 👍 2806 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/*
	解答成功:
				执行耗时:4 ms,击败了96.73% 的Go用户
				内存消耗:3.2 MB,击败了100.00% 的Go用户

	解题思路：遍历数组，当前子切片的最大和可能来源于，当前子切片的累加和，也有可能前面切片的累加值。
	而前面切片切片的累加和的最大值有可能是最前面的元素累加的和，然后加上当前元素的和，也有可能是当前元素的和。
 */


func maxSubArray(nums []int) int {
	accum,max := 0,nums[0]

	for _,val :=range nums {
		accum = maxVal(accum+val,val)
		max = maxVal(accum,max)
	}
	return max
}

func maxVal(a,b int)int{
	if a > b{
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)


