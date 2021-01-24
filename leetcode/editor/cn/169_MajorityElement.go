//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。 
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。 
//
// 
//
// 示例 1： 
//
// 
//输入：[3,2,3]
//输出：3 
//
// 示例 2： 
//
// 
//输入：[2,2,1,1,1,2,2]
//输出：2
// 
//
// 
//
// 进阶： 
//
// 
// 尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。 
// 
// Related Topics 位运算 数组 分治算法 
// 👍 840 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
package main
/*
	方法一:遍历数组,使用map去统计每个数的次数,最后取出次数最多的元素
    解答成功:
			执行耗时:20 ms,击败了84.74% 的Go用户
			内存消耗:6.1 MB,击败了23.44% 的Go用户
	时间复杂度:O(n)  空间复杂度：O(n)
*/

//func majorityElement(nums []int) int {
//	numMap :=make(map[int]int,0)
//	for _,num :=range nums{
//		val := numMap[num]
//		numMap[num] = val+1
//	}
//	// 遍历map，找出最大值对应的key
//
//	var count,key  = 0,0
//
//	for k,v :=range numMap{
//		if v > count {
//			count = v
//
//			key = k
//		}
//	}
//	return key
//}
/*
	对上面的方法进行优化:使用map统计数字出现的次数，当次数大于⌊ n/2 ⌋时，我们返回这个数字，返回结束遍历
    解答成功:
						执行耗时:16 ms,击败了98.94% 的Go用户
						内存消耗:6.1 MB,击败了23.44% 的Go用户
*/
//func majorityElement(nums []int) int {
//	numMap :=make(map[int]int,0)
//	length:= len(nums)
//	for _,num :=range nums{
//		val := numMap[num]
//
//		if val >= length/2 {
//			return  num
//		}
//		numMap[num] = val+1
//	}
//	// 由于出现最多次数的元素存在，所以肯定不会走到这，写在这块的代码是为了可以编译通过
//	return -1
//}


/*
    //而更巧妙的办法是使用球类运动中的盯人战术，只要两个选票数字不一样，就相互盯死，
    // 那么最后剩下的没被盯住肯定是那个超过半数的选票，这个就是摩尔投票法

	时间和上一个差不多，但是空间复杂度要更小
	解答成功:
						执行耗时:16 ms,击败了98.94% 的Go用户
						内存消耗:6.1 MB,击败了76.74% 的Go用户
*/
//
func majorityElement(nums []int) int {
	result := nums[0]
	times := 1

	for i,length:=1,len(nums);i<length;i++{
		num := nums[i]
		if num == result {
			times++
		}else{
			times--
		}

		if times == -1{
			result = num
			times = 1
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)


