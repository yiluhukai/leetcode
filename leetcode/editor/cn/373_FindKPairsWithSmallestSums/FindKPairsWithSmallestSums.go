package main

import (
	"container/heap"
)

//给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。
//
// 找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。
//
// 示例 1:
//
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//输出: [1,2],[1,4],[1,6]
//解释: 返回序列中的前 3 对数：
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
//
// 示例 2:
//
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//输出: [1,1],[1,1]
//解释: 返回序列中的前 2 对数：
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
//
// 示例 3:
//
// 输入: nums1 = [1,2], nums2 = [3], k = 3
//输出: [1,3],[2,3]
//解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
// Related Topics 数组 堆（优先队列）
// 👍 193 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
type pairs [][]int

func (p pairs) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p pairs) Less(i, j int) bool {
	return p[i][0]+p[i][1] < p[j][0]+p[j][1]
}

// Swap swaps the elements with indexes i and j.
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//解答成功:
//执行耗时:108 ms,击败了36.23% 的Go用户
//内存消耗:20 MB,击败了50.72% 的Go用户
// 时间复杂度：N = n1* n2  O(NlogN)
// 空间复杂度：O(n)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	length1, length2 := len(nums1), len(nums2)
	length := length1 * length2
	allPairs := make([][]int, 0, length)

	for _, x := range nums1 {
		for _, y := range nums2 {
			pair := []int{x, y}
			allPairs = append(allPairs, pair)
		}
	}
	// 对上面的所有组合进行排序
	sort.Sort(pairs(allPairs))

	// 取出前k个组合
	count := min(k, length)

	return allPairs[:count]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
*/


// 使用最大堆来完成这个工作:定义一个长度未k的最大堆，
// 遍历所有的组合，当堆的长度小于k的时候或当前组合小于堆顶元素时，若堆的长度未k，删除堆顶元素再将当前组合插入，
// 否则直接插入当前组合
//执行耗时:4 ms,击败了98.59% 的Go用户
//内存消耗:4.1 MB,击败了71.83% 的Go用户
// 时间复杂度为

//type heapPair [][]int
//
//
//func (p heapPair) Len() int {
//	return len(p)
//}
//
//// Less reports whether the element with
//// index i should sort before the element with index j.
//func (p heapPair) Less(i, j int) bool {
//	return p[i][0]+p[i][1] > p[j][0]+p[j][1]
//}
//
//// Swap swaps the elements with indexes i and j.
//func (p heapPair) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//// add x as element Len()
//func (p *heapPair) Push(x interface{}){
//	*p= append(*p, x.([]int))
//}
//// remove and return element Len() - 1.
//func (p *heapPair)Pop() interface{}{
//	length := len(*p)
//	if length < 1 {
//		return nil
//	}else{
//		v := (*p)[length-1]
//		*p =(*p)[:length-1]
//		return v
//	}
//}
//
//// return  element Len()-1
//func (p heapPair)Peek()interface{}{
//	length := len(p)
//	if length < 1 {
//		return nil
//	}
//	return p[0]
//}
//
//func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
//	// 定义一个最大堆
//	h:=&heapPair{}
//	heap.Init(h)
//	length1,length2 := len(nums1),len(nums2)
//	//数组是升序排列的，所以最小的k个组合一定不会包含超出k的部分
//	for i:=0;i<length1;i++ {
//		for j:=0;j<length2;j++{
//			if  l := h.Len(); l < k || nums1[i]+ nums2[j] < h.Peek().([]int)[0] + h.Peek().([]int)[1]{
//				if l == k{
//					// 删除堆顶的元素
//					heap.Pop(h)
//				}
//				// 插入当前的组合
//				heap.Push(h,[]int{nums1[i],nums2[j]})
//			}
//		}
//	}
//	// 将堆中的所有元素取出返回
//	//fmt.Printf("%#v\n",*h)
//	return *h
//}
//
//func min(a, b int) int {
//	if b < a {
//		return b
//	}
//	return a
//}



// 基于小顶堆的最优实现:
// 假设输入:nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 那么所有的组合会形成一个二维数组
//    1,2  1,4  1,6
//    7,2  7,4  7,6
//    11,2 11,4 11,6
//  当我们取最小的k个组合，k个组合一定是k*k个组合中产生的。
//  我们可以看出来二维数组的每一行是递增，那么我们将二维数组第一列中(每行的最小值)全部或者k个全部放入最小堆中
//  每次从堆中去堆顶元素，再将该行的下一个元素放入堆中，直到取够k个元素或者堆为空
//  为了方便去查找堆中的一个元素，我们放入堆中组合的下标

//3:11 下午	info
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:3.9 MB,击败了77.78% 的Go用户
type heapPair struct {
	// 存放堆中元素(组合的下标)的位置
	arr [][]int
	// 存放数组的地方,为了在Less方法中使用
	nums1,nums2 []int
}



func (p heapPair) Len() int {
	return len(p.arr)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p heapPair) Less(i, j int) bool {
	// 堆空间中相邻两个元素的比较
	before :=p.arr[i]
	after := p.arr[j]
	return p.nums1[before[0]]+p.nums2[before[1]] < p.nums1[after[0]]+ p.nums2[after[1]]
}

// Swap swaps the elements with indexes i and j.
func (p heapPair) Swap(i, j int) {
	p.arr[i], p.arr[j] = p.arr[j], p.arr[i]
}
// add x as element Len()
func (p *heapPair) Push(x interface{}){
	p.arr= append(p.arr, x.([]int))
}
// remove and return element Len() - 1.
func (p *heapPair)Pop() interface{}{
	length := p.Len()
	if length < 1 {
		return nil
	}else{
		v := (p.arr)[length-1]
		p.arr =(p.arr)[:length-1]
		return v
	}
}

// return  element Len()-1
func (p heapPair)Peek()interface{}{
	length := p.Len()
	if length < 1 {
		return nil
	}
	return p.arr[0]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 定义一个最大堆
	h:=&heapPair{
		nums1: nums1,
		nums2: nums2,
	}
	heap.Init(h)
	length1,length2 := len(nums1),len(nums2)
	// 当有一个数组为空的时候，这个时候返回的肯定是一个空切片
	if length1 == 0  || length2 == 0{
		return nil
	}
	// 数组是升序排列的，最小的一定是下标 0 0的组合
	// 我们选取包含数组num1的组合放入最小堆中，每次出去堆顶的元素
	// 如果存在比当前元素大的
	for i:=0;i<min(length1,k);i++ {
		pair := []int{ i,0}
		heap.Push(h,pair)
	}
	res := make([][]int,0,k)

	//fmt.Printf("%#v\n",h)
	//
	for k >0 && h.Len() > 0 {
		//取出当前的堆顶元素
		top := heap.Pop(h).([]int)
		// 将堆顶的元素加入结果
		res = append(res, []int{nums1[top[0]],nums2[top[1]]})

		if top[1] < length2-1 {
			heap.Push(h,[]int{ top[0], top[1]+1})
		}
		//fmt.Printf("%#v\n", h)
		k--
	}
	return res
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
