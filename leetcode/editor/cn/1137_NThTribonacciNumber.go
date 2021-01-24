//泰波那契序列 Tn 定义如下： 
//
// T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2 
//
// 给你整数 n，请返回第 n 个泰波那契数 Tn 的值。 
//
// 
//
// 示例 1： 
//
// 输入：n = 4
//输出：4
//解释：
//T_3 = 0 + 1 + 1 = 2
//T_4 = 1 + 1 + 2 = 4
// 
//
// 示例 2： 
//
// 输入：n = 25
//输出：1389537
// 
//
// 
//
// 提示： 
//
// 
// 0 <= n <= 37 
// 答案保证是一个 32 位整数，即 answer <= 2^31 - 1。 
// 
// Related Topics 递归 
// 👍 54 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:1.9 MB,击败了81.67% 的Go用户
var cache = make(map[int]int,35)
func tribonacci(n int) int {

	if val,ok := cache[n];ok{
		return val
	}

	if n == 0{
		cache[0] =0
		return  0
	}
	if n == 1 || n==2{
		cache[n] = 1
		return 1
	}

	res := tribonacci(n-1)+tribonacci(n-2)+tribonacci(n-3)

	cache[n] = res

	return res
}
//leetcode submit region end(Prohibit modification and deletion)


