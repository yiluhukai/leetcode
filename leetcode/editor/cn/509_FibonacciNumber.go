//斐波那契数，通常用 F(n) 表示，形成的序列称为斐波那契数列。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是： 
//
// F(0) = 0,   F(1) = 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
// 
//
// 给定 N，计算 F(N)。 
//
// 
//
// 示例 1： 
//
// 输入：2
//输出：1
//解释：F(2) = F(1) + F(0) = 1 + 0 = 1.
// 
//
// 示例 2： 
//
// 输入：3
//输出：2
//解释：F(3) = F(2) + F(1) = 1 + 1 = 2.
// 
//
// 示例 3： 
//
// 输入：4
//输出：3
//解释：F(4) = F(3) + F(2) = 2 + 1 = 3.
// 
//
// 
//
// 提示： 
//
// 
// 0 ≤ N ≤ 30 
// 
// Related Topics 数组 
// 👍 184 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)

// 循环

//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:1.9 MB,击败了98.52% 的Go用户
//func fib(n int) int {
//	if n <= 1 {
//		return n
//	}
//	// 递推公式
//	n1,n2 :=  0,1
//	for i:=2;i < n;i++{
//		n1,n2 = n2, n1+n2
//	}
//	return n1+n2
//}
//  递归
func fib(n int) int {
	//  递归的终止条件
	if n <= 1 {
		return n
	}
	// 递推公式
	return  fib(n-1) + fib(n-2)
}
//leetcode submit region end(Prohibit modification and deletion)


