package  main
//请完成一个函数，输入一个二叉树，该函数输出它的镜像。 
//
// 例如输入： 
//
// 4 
// / \ 
// 2 7 
// / \ / \ 
//1 3 6 9 
//镜像输出： 
//
// 4 
// / \ 
// 7 2 
// / \ / \ 
//9 6 3 1 
//
// 
//
// 示例 1： 
//
// 输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
// 
//
// 
//
// 限制： 
//
// 0 <= 节点个数 <= 1000 
//
// 注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/ 
// Related Topics 树 
// 👍 127 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:2.1 MB,击败了18.13% 的Go用户
func mirrorTree(root *TreeNode) *TreeNode {
	if root ==nil {
		return nil
	}
	// 非空树
	// 交换当前节点左右子节点
	root.Left,root.Right = root.Right,root.Left
	// 对子树生成镜像
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
