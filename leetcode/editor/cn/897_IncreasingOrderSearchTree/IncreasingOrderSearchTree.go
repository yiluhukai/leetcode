package  main

//给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
//
// 
//
// 示例 1： 
//
// 
//输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
//输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
// 
//
// 示例 2： 
//
// 
//输入：root = [5,1,7]
//输出：[1,null,5,null,7]
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数的取值范围是 [1, 100] 
// 0 <= Node.val <= 1000 
// 
// Related Topics 树 深度优先搜索 递归 
// 👍 147 👎 0
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.2 MB,击败了48.30% 的Go用户
//func increasingBST(root *TreeNode) *TreeNode {
//	if root ==nil {
//		return nil
//	}
//	// 中序遍历 莫里斯遍历
//	res := make([]int,0)
//	for root != nil {
//		if root.Left !=nil {
//			// 查找左子树的最右子树
//			r := root.Left
//			for r.Right != nil && r.Right!= root {
//				r = r.Right
//			}
//			if r.Right == nil {
//				r.Right = root
//				root = root.Left
//			}else{
//				// r.Right!= root
//				// root节点的左子树已经遍历完成了
//				r.Right = nil
//				res = append(res, root.Val)
//				root = root.Right
//			}
//		}else{
//			// 当前节点没有左子树
//			res = append(res, root.Val)
//			root = root.Right
//		}
//	}
//	// 对res 进行排序
//	sort.Ints(res)
//	head := new(TreeNode)
//	cur := head
//	length := len(res)
//	for index, v := range res {
//		cur.Val = v
//		if index != length-1 {
//			cur.Right = new(TreeNode)
//			cur = cur.Right
//		}
//	}
//	return head
//}
// 对原树作修改
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head := root
	// 递归去修改树的结构
	return head
}
//leetcode submit region end(Prohibit modification and deletion)
