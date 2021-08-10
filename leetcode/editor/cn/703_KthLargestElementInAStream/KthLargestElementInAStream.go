package  main

import (
	"container/heap"
)



//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
//
// 请实现 KthLargest 类： 
//
// 
// KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。 
// int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["KthLargest", "add", "add", "add", "add", "add"]
//[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
//输出：
//[null, 4, 5, 5, 8, 8]
//
//解释：
//KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
//kthLargest.add(3);   // return 4
//kthLargest.add(5);   // return 5
//kthLargest.add(10);  // return 5
//kthLargest.add(9);   // return 8
//kthLargest.add(4);   // return 8
// 
//
// 
//提示：
//
// 
// 1 <= k <= 104 
// 0 <= nums.length <= 104 
// -104 <= nums[i] <= 104 
// -104 <= val <= 104 
// 最多调用 add 方法 104 次 
// 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素 
// 
// Related Topics 树 设计 二叉搜索树 二叉树 数据流 堆（优先队列） 
// 👍 281 👎 0


//leetcode submit region begin(Prohibit modification and deletion)

// 使用容量为k小顶堆来实现，当使用add方法时比较当前堆顶元素和插入元素，
// 当堆顶的元素大于当前元素时，替换，否则直接返回堆顶的元素
type HeapInt []int


// Len is the number of elements in the collection.
func (h HeapInt)Len() int{
	return len(h)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (h HeapInt)Less(i, j int) bool{
	return	h[i] < h[j]
}
// Swap swaps the elements with indexes i and j.
func (h HeapInt)Swap(i, j int){
	h[i],h[j] =  h[j],h[i]
}

// add x as element Len()
func (h *HeapInt)Push(x interface{}) {
	*h = append(*h, x.(int))
}
// remove and return element Len() - 1.
func (h *HeapInt)Pop() interface{} {
	length := h.Len()
	v := (*h)[length-1]
	*h = (*h)[:length-1]
	return v
}
// 获取堆顶的元素
func (h HeapInt)Peek()int{
	return h[0]
}
type KthLargest struct {
	// 最小堆
	nums *HeapInt
	capacity int
}
// 实现堆上


func Constructor(k int, nums []int) KthLargest {
	//执行耗时:28 ms,击败了97.81% 的Go用户
	//内存消耗:7.8 MB,击败了91.80% 的Go用户
	//sort.Ints(nums)
	length :=len(nums)
	//h := HeapInt(nums[length-min(length,k):])

	//		解答成功:
	//					执行耗时:24 ms,击败了100.00% 的Go用户
	//					内存消耗:7.8 MB,击败了91.80% 的Go用户
	h := HeapInt(nums[:min(length,k)])
	heap.Init(&h)
	// 向堆中加入元素，需要维持堆的容量不超过k
	for i:=k;i<length;i++{
		if v:=nums[i];h.Peek() < v{
			heap.Pop(&h)
			heap.Push(&h,v)
		}
	}


	return KthLargest{
		nums: &h,
		capacity: k,
	}
}

func (this *KthLargest) Add(val int) int {
	// 堆中有k个元素 且 堆顶元素小于 val
	if this.capacity == this.nums.Len() && val > this.nums.Peek(){
			heap.Pop(this.nums)
	}
	//fmt.Printf("nums=%#v\n",this.nums)
	if this.capacity > this.nums.Len(){
		heap.Push(this.nums,val)
	}
	//fmt.Printf("val=%v,nums=%#v\n",val,this.nums)
	if this.capacity > this.nums.Len(){
		return 0
	}else{
		return this.nums.Peek()
	}
}



func min(a,b int)int{
	if b < a{
		return b
	}
	return a
}
//
//func main()  {
//	obj := Constructor(7, []int{ -10,1,3,1,4,10,3,9,4,5,1})
//	//-10,1,1,1,3,3,4,4,5,9,
//	v :=obj.Add(3)
//	fmt.Printf("%v\n",v)
//
//
//}
/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
//leetcode submit region end(Prohibit modification and deletion)


