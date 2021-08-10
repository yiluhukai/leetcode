package  main

import (
	"container/heap"
	"sort"
)

//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。 
//
// 
//
// 示例 1: 
//
// 
//输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
// 
//
// 示例 2: 
//
// 
//输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4 
//
// 
//
// 提示： 
//
// 
// 1 <= k <= nums.length <= 104 
// -104 <= nums[i] <= 104 
// 
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 
// 👍 1184 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
//思路一：先对数组排序，然后返回数组倒数第k个元素
//解答成功:
//执行耗时:8 ms,击败了82.10% 的Go用户
//内存消耗:3.4 MB,击败了100.00% 的Go用户
// 时间复杂度 O(NlogN)
//func findKthLargest(nums []int, k int) int {
//	sort.Ints(nums)
//	return nums[len(nums)-k]
//}

// 使用小顶堆来实现，设置堆的容量为k
// 当堆中的元素为k时，每次先和堆顶元素做比较，当大于堆顶元素，则删除堆顶元素
// 将该元素入堆，当所有元素都入堆，堆顶元素就是第K大的元素
// 时间复杂度：O(N) * LogK ==> O(N)
//解答成功:
//执行耗时:4 ms,击败了99.38% 的Go用户
//内存消耗:3.9 MB,击败了25.69% 的Go用户
type HeapInt struct {
	// 堆中存储数组的
	sort.IntSlice
	//	堆的最大长度
	size int
}

// remove and return element Len() - 1.
func (h *HeapInt)Pop()interface{}{
	length := len(h.IntSlice)
	v:=h.IntSlice[length-1]
	h.IntSlice = h.IntSlice[:length-1]
	return v
}
// add x as element Len()
func (h *HeapInt)Push(x interface{}){
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *HeapInt)Peek()int{
	return h.IntSlice[0]
}

func findKthLargest(nums []int, k int) int {
	h := &HeapInt{
		size:k,
	}
	heap.Init(h)

	for _,v :=range nums {
		if h.Len() == h.size && h.Peek() < v {
			heap.Pop(h)
		}
		if h.Len() < h.size{
			heap.Push(h,v)
		}
	}
	return h.Peek()
}

//leetcode submit region end(Prohibit modification and deletion)


