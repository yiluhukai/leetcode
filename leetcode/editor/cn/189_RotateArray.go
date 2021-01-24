//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。 
//
// 示例 1: 
//
// 输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
// 
//
// 示例 2: 
//
// 输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释: 
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100] 
//
// 说明: 
//
// 
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。 
// 要求使用空间复杂度为 O(1) 的 原地 算法。 
// 
// Related Topics 数组 
// 👍 758 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int)  {

	/*
		将数组的元素分为n-k和k两部分，先对这两部分进行反转，在对整个数组进行反转
	*/
	length:=len(nums)
	// 数组长度小于1，可以不用旋转
	if length <=1 || k ==0{
		return
	}

	// 当k大于数组的长度，相当于旋转 k%length次
	if k > length {
		k = k%length
	}

	reverseArr(nums,0,length-k)
	reverseArr(nums,length-k,k)
	reverseArr(nums,0,length)
}

func reverseArr(nums []int,start,length int){
	if length <=0 {
		return
	}
	for i:=start;i<length/2+start;i++{
		// 反序操作
		nums[i],nums[start+length-i-1+start] = nums[start+length-i-1+start],nums[i]
	}
}
//leetcode submit region end(Prohibit modification and deletion)


