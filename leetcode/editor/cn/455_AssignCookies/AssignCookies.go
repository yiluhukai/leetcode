package  main

import "sort"

//假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//
// 对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i
//]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。 
// 
//
// 示例 1: 
//
// 
//输入: g = [1,2,3], s = [1,1]
//输出: 1
//解释: 
//你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
//虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
//所以你应该输出1。
// 
//
// 示例 2: 
//
// 
//输入: g = [1,2], s = [1,2,3]
//输出: 2
//解释: 
//你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
//你拥有的饼干数量和尺寸都足以让所有孩子满足。
//所以你应该输出2.
// 
//
// 
//
// 提示： 
//
// 
// 1 <= g.length <= 3 * 104 
// 0 <= s.length <= 3 * 104 
// 1 <= g[i], s[j] <= 231 - 1 
// 
// Related Topics 贪心算法 
// 👍 295 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
	题目没有说胃口大小还有饼干的大小是有序的，所以先给他们排个序
	然后使用贪心算法：给每个孩子的饼干是满足他胃口的最小的饼干，如果存在这样的饼干，数量加一
	最后统计总的数量
	执行耗时:44 ms,击败了16.49% 的Go用户
	内存消耗:6.2 MB,击败了90.96% 的Go用户
	最坏的情况:有m块饼干n个小子，小孩子的胃口大于最大的饼干面积，此时的比较次数为m*n,所以时间复杂度为O(m*n)
	空间复杂度为O(1)
 */
func findContentChildren(g []int, s []int) int {
	// 排成从小到大的顺序

	sort.Ints(g)
	sort.Ints(s)
	count :=0

	for j,length:=0,len(s);j< length;j++ {
		pieSize:=s[j]
		kidsCount:=len(g)
		//当没有小孩子时提前终止
		if kidsCount ==0 {
			break
		}
		for i:=0;i<kidsCount;i++{
			//  遍历小孩子的胃口，找到满足这个小孩子胃口的最小的饼干
			if pieSize >= g[i] {
				count++
				g=g[i+1:]
				break
			}
		}
	}
	return count
}

/**

	使用双指针去代替上面的双层for循环
	执行耗时:36 ms,击败了86.87% 的Go用户
	内存消耗:6.2 MB,击败了90.96% 的Go用户

	时间复杂度为O(m),m为饼干的数量
	空间复杂度为O(1)
 */

//func findContentChildren(g []int, s []int) int {
//		// 排成从小到大的顺序
//		pieCount,kidsCcount := len(s),len(g)
//		if pieCount == 0 || kidsCcount ==0{
//			return 0
//		}
//		sort.Ints(g)
//		sort.Ints(s)
//
//		gi,si:=0,0
//
//		for gi < kidsCcount && si <pieCount {
//			if s[si] >= g[gi] {
//				gi++
//			}
//			si++
//		}
//		return gi
//}
//leetcode submit region end(Prohibit modification and deletion)


