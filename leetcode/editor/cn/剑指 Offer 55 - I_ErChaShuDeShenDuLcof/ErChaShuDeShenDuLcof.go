package main


//输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
//
// 例如：
//
// 给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
//
//
//
// 提示：
//
//
// 节点总数 <= 10000
//
//
// 注意：本题与主站 104 题相同：https://leetcode-cn.com/problems/maximum-depth-of-binary-tre
//e/
// Related Topics 树 深度优先搜索
// 👍 108 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BFS(广度优先遍历)
//解答成功:
//执行耗时:4 ms,击败了88.59% 的Go用户
//内存消耗:4.5 MB,击败了7.86% 的Go用户
// 时间复杂度O(n) 空间复杂度：O(n) n代表节点的个数
//func maxDepth(root *TreeNode) int {
//	var l = list.New()
//	if root == nil {
//		return 0
//	}
//	// 不为空
//	l.PushBack(root)
//	// 可以使用循环代替递归
//	//return  getDepth(l,0)
//	res := 0
//	for len := l.Len(); len != 0; len = l.Len() {
//		// 处理二叉树每一层的节点
//		for i := 0; i < len; i++ {
//			e := l.Remove(l.Front()).(*TreeNode)
//			if e.Left != nil {
//				l.PushBack(e.Left)
//			}
//			if e.Right != nil {
//				l.PushBack(e.Right)
//			}
//		}
//		res++
//	}
//	return res
//}
// 使用递归的方式
//func getDepth(l *list.List,n int)int{
//	len := l.Len()
//	if len == 0{
//		return n
//	}
//	for i:=0;i<len;i++{
//		e := l.Remove(l.Front()).(*TreeNode)
//		if e.Left !=nil{
//			l.PushBack(e.Left)
//		}
//		if e.Right !=nil{
//			l.PushBack(e.Right)
//		}
//	}
//	return getDepth(l,n+1)
//}
// DFS (深度优先遍历)
//

//执行耗时:4 ms,击败了88.59% 的Go用户
//内存消耗:4.2 MB,击败了88.05% 的Go用户

// 时间复杂度：O(n) 空间复杂度:O(h) n代表节点的个数 h代表节点的深度
// 最坏的情况下节点为n,树的深度为n,最好的情况下：floor(log2n)+1

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	// 返回当前节点的深度 + 左右子树中的最大深度
	return 1+max(maxDepth(root.Left),maxDepth(root.Right))
}

func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}



//leetcode submit region end(Prohibit modification and deletion)
