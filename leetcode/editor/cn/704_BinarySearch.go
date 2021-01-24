//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否
//则返回 -1。
//
//
//示例 1:
//
// 输入: nums = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 出现在 nums 中并且下标为 4
//
//
// 示例 2:
//
// 输入: nums = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不存在 nums 中因此返回 -1
//
//
//
//
// 提示：
//
//
// 你可以假设 nums 中的所有元素是不重复的。
// n 将在 [1, 10000]之间。
// nums 的每个元素都将在 [-9999, 9999]之间。
//
// Related Topics 二分查找
// 👍 189 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
//  循环
//			执行耗时:40 ms,击败了82.18% 的Go用户
//						内存消耗:6.7 MB,击败了9.54% 的Go用户
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	start, end := 0, length-1

	for start <= end {
		if half := (end + start) / 2; nums[half] == target {
			return half
		} else if nums[half] < target {
			//右半边去查找
			start = half + 1
		} else {
			//左半边去查找
			end = half - 1
		}
	}
	return -1
}
//递归实现
//执行耗时:40 ms,击败了82.18% 的Go用户
//内存消耗:6.7 MB,击败了35.63% 的Go用户
//func search(nums []int, target int) int{
//	return binaryFind(nums,target,0,len(nums))
//}
//
//func binaryFind(nums []int,target,start,end int)int{
//	length := len(nums[start:end])
//	if length == 0{
//		return  -1
//	}
//	// 递归的过程和递推条件
//	if half:=(end+start)/2;nums[half] == target{
//		return half
//	}else if nums[half] < target{
//		//右半边去查找
//		return binaryFind(nums,target,half+1,end)
//	}else{
//		//左半边去查找
//		return binaryFind(nums,target,start,half)
//	}
//}
//leetcode submit region end(Prohibit modification and deletion)
