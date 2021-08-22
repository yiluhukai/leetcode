package main

//如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
//
//
// n >= 3
// 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
//
//
// 给定一个严格递增的正整数数组形成序列 arr ，找到 arr 中最长的斐波那契式的子序列的长度。如果一个不存在，返回 0 。
//
// （回想一下，子序列是从原序列 arr 中派生出来的，它从 arr 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8]
//是 [3, 4, 5, 6, 7, 8] 的一个子序列）
//
//
//
//
//
//
// 示例 1：
//
//
//输入: arr = [1,2,3,4,5,6,7,8]
//输出: 5
//解释: 最长的斐波那契式子序列为 [1,2,3,5,8] 。
//
//
// 示例 2：
//
//
//输入: arr = [1,3,7,11,12,14,18]
//输出: 3
//解释: 最长的斐波那契式子序列有 [1,11,12]、[3,11,14] 以及 [7,11,18] 。
//
//
//
//
// 提示：
//
//
// 3 <= arr.length <= 1000
//
// 1 <= arr[i] < arr[i + 1] <= 10^9
//
//
//
//
//
// 注意：本题与主站 873 题相同： https://leetcode-cn.com/problems/length-of-longest-fibonacc
//i-subsequence/
// Related Topics 数组 哈希表 动态规划
// 👍 2 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 暴力搜索：
// 斐波那契式子序列的长度大于3，且X_i + X_{i+1} = X_{i+2}
// 对正整数数组中的所有中任意两个数作为斐波那契式子序列中的前两个元素，然后得出后续的节点，并在数组中查询它的出现次数
//解答成功:
//执行耗时:184 ms,击败了100.00% 的Go用户
//内存消耗:4.7 MB,击败了85.71% 的Go用户

// 时间复杂度为 O(N^2 * log M) N代表了正整数数组的长度，M代表了数组中的最大值
// 空间复杂度为：O(N)
/*
func lenLongestFibSubseq(arr []int) int {

	length := len(arr)
	// 为了快速从数组中查找数字是否存在，可以使用map来快速查找
	mp := make(map[int]bool, length)
	for _, num := range arr {
		mp[num] = true
	}
	maxLength := 0
	// 从arr中取出两个数作为斐波那契式子序列的前两个元素
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			x, y := arr[i], arr[j]
			res := 2
			for mp[x+y] {
				res++
				x, y = y, x+y
			}

			if res >= 3 {
				maxLength = max(res, maxLength)
			}
		}
	}
	return maxLength
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

// 动态规划：
// 1. 原问题找到 arr 中最长的斐波那契式的子序列的长度 子问题：arr[i],arr[j] 斐波那契式的子序列中最后出现的元素是子序列的长度
// 2. 设置状态 Dp[(i,j)] = 斐波那契式的子序列中最后出现的元素是 arr[i],arr[j]的子序列的长度
// 3. 状态转移方程 arr[ i] + arr[j] = arr[k] ,那么斐波那契式的子序列为arr[i],arr[j],arr[k],转化为节点表示法为(i,j),(j,k)
// Dp(j,k)代表以arr[j],arr[k]为结尾的斐波那契式的子序列的长度 Dp[i,j] + 1 = Dp[j,k]
// 我们的目标值是求所有 Dp[i,j]中的最大值
// 4. 处理边界 Dp[(i,j]+ 1 = Dp[j,k] , i >=0 ,Dp[i,j]的最小长度为2

//解答成功:
//执行耗时:180 ms,击败了100.00% 的Go用户
//内存消耗:6.3 MB,击败了85.71% 的Go用户
// 时间复杂度为O(N*N)

func lenLongestFibSubseq(arr []int) int {
	length := len(arr)
	// arr中值到 index的映射
	index := make(map[int]int,length)
	for i,v :=range arr{
		index[v] = i
	}
	//x,y => dp[x,y]的映射,这里为了可以根据(x,y)作为map的key,我们将 （x,y）转换成x*length+y
	dpArr := make(map[int]int,length)
	maxLength := 0
	// arr[i] + arr[j]  = arr [k]
	// 我们要求所有的dp[j,k]，那么我们先要去dp[i,j] ; 然后dp[j,k] =dp[i,j]+1
	// j,k可以是数组中的任意两个

	for  j :=0;j < length-1;j++{
		for k := j+1;k <length;k++ {
			//求a[i]的中i的值
			num :=  arr[k] - arr[j]
			// 利用mp快速查找arr[i]的下标
			i :=-1
			v,ok := index[num]
			if ok {
				i = v
			}
			if i >= 0 && i < j {
				dp,ok := dpArr[i*length+j]
				// 最小的dp[i,j]
				if !ok {
					dp= 2
				}
				// 根据dp[i,j]去求dp[j,k]
				key :=  j*length+k
				dpArr[key] = dp + 1
				if dpArr[key] >= 3 {
					maxLength = max( dpArr[key], maxLength)
				}
			}
		}
	}
	return maxLength
}

func max(a,b int)int  {
	if a > b{
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
