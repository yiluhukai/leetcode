package  main
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。 
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？ 
//
// 注意：给定 n 是一个正整数。 
//
// 示例 1： 
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶 
//
// 示例 2： 
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
// 
// Related Topics 记忆化搜索 数学 动态规划 
// 👍 1806 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 递归：n阶阶梯：每次只能爬一步或者两步：当n = 1是，只有一种爬发 1阶，当n = 2时，有两种 1阶 + 1阶 或者 2阶
// 当n > 2时，最后一步可以爬一阶也可以爬二阶； 那么 C(n) = C(n-1) + C(n-2)其中C（n）代表爬n阶的爬法
// climbStairs(n-1)时需要计算climbStairs(n-2)到climbStairs(2)
// climbStairs(n-2)时需要计算climbStairs(n-3)到到climbStairs(2)
// 为了减少计算次数，我们将结果缓存起来

//	解答成功:
//					执行耗时:0 ms,击败了100.00% 的Go用户
//					内存消耗:2 MB,击败了5.43% 的Go用户
//var times = make(map[int]int)
//func climbStairs(n int) int {
//	if v,ok := times[n]; ok {
//		return v
//	}
//	if  n  <= 2{
//		return n
//	}else{
//		time := climbStairs(n-1) + climbStairs(n-2)
//		times[n] = time
//		return time
//	}
//}

// 动态规划
// 1。原问题和子问题 原问题：climbStairs(n) 子问题climbStairs(i) i = 1 ... n
// 2. 设计状态  DP[i]代表爬n阶台阶的爬法
// 3。状态转移方程 DP[i] = DP[i-1] + DP[i-2] i>=3
// 4. 确定边界状态值  DP[1] = 1;DP[2] = 2
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:1.9 MB,击败了18.77% 的Go用户
func climbStairs(n int) int {
  dp:=make([]int,n+1)
  for i:=1;i <= n;i++{
  	if  i <= 2 {
  		dp[i] = i
	}else{
		dp[i] = dp[i-1] + dp[i-2]
	}
  }
  return dp[n]
}
//leetcode submit region end(Prohibit modification and deletion)


