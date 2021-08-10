package  main

//给你一个整数数组 target 和一个数组 initial ，initial 数组与 target 数组有同样的维度，且一开始全部为 0 。
//
// 请你返回从 initial 得到 target 的最少操作次数，每次操作需遵循以下规则： 
//
// 
// 在 initial 中选择 任意 子数组，并将子数组中每个元素增加 1 。 
// 
//
// 答案保证在 32 位有符号整数以内。 
//
// 
//
// 示例 1： 
//
// 输入：target = [1,2,3,2,1]
//输出：3
//解释：我们需要至少 3 次操作从 intial 数组得到 target 数组。
//[0,0,0,0,0] 将下标为 0 到 4 的元素（包含二者）加 1 。
//[1,1,1,1,1] 将下标为 1 到 3 的元素（包含二者）加 1 。
//[1,2,2,2,1] 将下表为 2 的元素增加 1 。
//[1,2,3,2,1] 得到了目标数组。
// 
//
// 示例 2： 
//
// 输入：target = [3,1,1,2]
//输出：4
//解释：(initial)[0,0,0,0] -> [1,1,1,1] -> [1,1,1,2] -> [2,1,1,2] -> [3,1,1,2] (tar
//get) 。
// 
//
// 示例 3： 
//
// 输入：target = [3,1,5,4,2]
//输出：7
//解释：(initial)[0,0,0,0,0] -> [1,1,1,1,1] -> [2,1,1,1,1] -> [3,1,1,1,1] 
//                                  -> [3,1,2,2,2] -> [3,1,3,3,2] -> [3,1,4,4,2]
// -> [3,1,5,4,2] (target)。
// 
//
// 示例 4： 
//
// 输入：target = [1,1,1,1]
//输出：1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= target.length <= 10^5 
// 1 <= target[i] <= 10^5 
// 
// Related Topics 栈 贪心 数组 动态规划 单调栈 
// 👍 37 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 线段树 + 分治
// 找到target中当前区域[0,n-1]的最小值target[index]，将当前值s(上次增加的后的值，默认为0)变为target[index]，记录val = val + target[index]-s
// 然后当前区域划分成[0,index-1] 和 [index+1,n-1]两个区域继续上面额度过程
// 线段树的作用是为了让我们更好的求当前区域的最小值，我们在线段树中存储的最小值在线段树中的下标

// 创建一颗线段树

//解答成功:
//执行耗时:248 ms,击败了12.50% 的Go用户
//内存消耗:25.5 MB,击败了12.50% 的Go用户
// 时间复杂度 数组的区间长度为N,那么生成的线段的长度为4*N,我们是用递归来构造线段树的节点的，那么时间复杂度为O(log 4*N)
// 对于N个区间，我们可能需要查找每个区间的最小值，每个区间的最小值的查找时间复杂度为log(4 * N)
// 所以最终的时间复杂度为 O(N*Log 4* N) +O(N) ==> O(NlogN)

// 空间复杂度为: O(log 4*N) == > O(log N)

//type SegmentTree struct {
//	// 初始化传入的数组
//	nums []int
//	// 使用build方法构造的线段树
//	values []int
//	size int// 线段树的长度
//	// 原始区间的宽度
//	n int
//}
//// 因为线段树是一个完全二叉树，且叶子节点的数目等于原数组的长度 length，
//// 假设线段树的最后一个节点的下标为 n,那么线段树中第一个非叶子节点的下标为,当 n是偶数的时候(n-2)/2，当n为奇数的时候 (n-1)/2
//// 那么线段树中的叶子节点的个数 n - (n-1)/2 = length 或者 n - (n-2)/2 = length
////    n = 2length-1  或者 n == 2length-2
//// 	长度为length的元素组可以得到一个长度为n+1(2length || 2length-1)的线段树。
//
//func NewSegmentTree(nums []int)*SegmentTree{
//	length := len(nums)
//	s := &SegmentTree{
//		nums:nums,
//		values: make([]int, 4 * length),
//		n: length-1,
//	}
//	s.build(0,0,s.n)
//	return s
//}
//// s.values中的下标，原始的区间长度[left,right]
//// values中保存的是nums最小值的下标
//func (s * SegmentTree) build(pos,left,right int){
//	// 没递归一层，线段树的长度就加1
//	s.size++
//	if left == right {
//		s.values[pos] = left
//		return
//	}
//	mid := (left + right)/2
//
//	s.build(pos*2+1,left,mid)
//	s.build(pos*2+2,mid+1,right)
//	leftIndex,rightIndex := s.values[pos*2+1],s.values[pos*2+2]
//	if s.nums[leftIndex] > s.nums[rightIndex] {
//		s.values[pos] = rightIndex
//		return
//	}
//	s.values[pos] = leftIndex
//}
//
//// 查询区间[qLeft,qRight]中的最小值在nums中的下标
//func (s * SegmentTree)getMinValue(pos,qLeft,qRight,left,right int)int{
//	if qLeft > qRight{
//		return -1
//	}
//	if qLeft == left && qRight == right{
//		return s.values[pos]
//	}
//	// 子节点的区间[left,mid] [mid+1,right]
//	// 将区间[qLeft,qRight]拆分成[qLeft,min(qRight,mid)] 和max[mid+1,qLeft,qRight]
//	mid := (left +right)/2
//
//	leftIndex := s.getMinValue(pos * 2 + 1,qLeft,min(qRight,mid) ,left,mid)
//	rightIndex := s.getMinValue(pos * 2 + 2,max(mid+1,qLeft),qRight,mid+1,right)
//
//
//	if leftIndex == -1 {
//		return rightIndex
//	}
//	if rightIndex == -1{
//		return leftIndex
//	}
//
//	if s.nums[leftIndex] > s.nums[rightIndex] {
//		return rightIndex
//	}
//	return  leftIndex
//}
//// 计算minValue 变为区间最小值的次数
//func(s SegmentTree)getMinNumberOperations(minValue,left,right int)int{
//	// 我们将区间划分成了[left,index -1] [index+1,right]
//	if left  >  right{
//		return 0
//	}
//	index := s.getMinValue(0,left,right, 0 ,s.n)
//	leftTimes := s.getMinNumberOperations(s.nums[index],left,index-1)
//	rightTimes := s.getMinNumberOperations(s.nums[index],index + 1,right)
//	return s.nums[index] - minValue + leftTimes + rightTimes
//}
//
//func minNumberOperations(target []int) int {
//	s := NewSegmentTree(target)
//	return s.getMinNumberOperations(0,0,s.n)
//}
//
//
//
//func min(a,b int)int{
//	if b < a {
//		return b
//	}
//	return a
//}
//
//func max(a,b int)int{
//	if b > a {
//		return b
//	}
//	return a
//}

// 使用差分数组求解
//解答成功:
//执行耗时:128 ms,击败了100.00% 的Go用户
//内存消耗:7.8 MB,击败了100.00% 的Go用户
// 时间复杂度为O(n)
func minNumberOperations(target []int) int {
	length:=len(target)
	if length == 0{
		return 0
	}
	total := target[0]
	for i:=1;i < length;i++{
		total += max(target[i]-target[i-1],0)
	}
	return total
}

func max(a,b int)int{
	if b > a {
		return b
	}
	return a
}
// 我们对大区间的数字整体加1的话，会使小区间的最小值的初始值发生变化
//func getMinNumberOperations(minVal int,)
//leetcode submit region end(Prohibit modification and deletion)


