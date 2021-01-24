//给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。 
//
// （例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。） 
//
// 返回 A 中好子数组的数目。 
//
// 
//
// 示例 1： 
//
// 输入：A = [1,2,1,2,3], K = 2
//输出：7
//解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
// 
//
// 示例 2： 
//
// 输入：A = [1,2,1,3,4], K = 3
//输出：3
//解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].
// 
//
// 
//
// 提示： 
//
// 
// 1 <= A.length <= 20000 
// 1 <= A[i] <= A.length 
// 1 <= K <= A.length 
// 
// Related Topics 哈希表 双指针 Sliding Window 
// 👍 135 👎 0
package main
/*
	total := 0
	for left,length := 0,len(A);left < length; left++ {
		count := 0
		mp := make(map[int]bool,0)
		for right:=left;right <length && count <= K;right++{
			if _,ok := mp[A[right]];!ok{
				mp[A[right]] = true
				count++
			}
			if count == K {
				total++
			}
		}
	}
	return total
*/
//解答成功:
//执行耗时:76 ms,击败了78.26% 的Go用户
//内存消耗:7 MB,击败了59.09% 的Go用户
//leetcode submit region begin(Prohibit modification and deletion)
func subarraysWithKDistinct(A []int, K int) int {
	//滑动窗口
	left,right,count,result:=0,0,0,0
	mp:= make(map[int]int)
	length:=len(A)
	for right < length{
		//  将right位置的数字的出现的次数放入mp中
		val:=mp[A[right]]
		//  第一次出现的时候val = 0
		mp[A[right]] = val + 1

		if val+1 == 1 {
			// 第一次出现
			count++
		}
		//  当count >k的时候，移动left
		for count > K {
			mp[A[left]] = mp[A[left]] - 1
			if mp[A[left]] == 0 {
				count --
			}
			left++
		}
		left2:=left
		// 当 count = k,滑动left2
		for count == K{
			result++
			//  移除mp中left2对应位置的数字-1
			val:=mp[A[left2]]
			if val-1 > 0{
				mp[A[left2]] = val-1
				left2++
			}else{
				break
			}
		}
		// 在mp中补会已经left2移动过程中删除的元素次数
		for i:=left;i<left2;i++{
			val := mp[A[i]]
			mp[A[i]] = val + 1
		}
		right++
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)


