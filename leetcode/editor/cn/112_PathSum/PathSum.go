package  main
//给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和
// targetSum 。 
//
// 叶子节点 是指没有子节点的节点。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：root = [1,2,3], targetSum = 5
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：root = [1,2], targetSum = 0
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// 树中节点的数目在范围 [0, 5000] 内 
// -1000 <= Node.val <= 1000 
// -1000 <= targetSum <= 1000 
// 
// Related Topics 树 深度优先搜索 
// 👍 571 👎 0

//Definition for a binary tree node.
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
//执行耗时:8 ms,击败了73.38% 的Go用户
//内存消耗:5.8 MB,击败了5.87% 的Go用户
//解答成功:
//执行耗时:8 ms,击败了73.38% 的Go用户
//内存消耗:4.6 MB,击败了100.00% 的Go用户
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return  hasPathSum(root.Left, targetSum - root.Val) ||  hasPathSum(root.Right, targetSum - root.Val)


	// 求出所有的路径和再去查找
	//result := getAllPathSum(root)
	//for _,v :=range result {
	//	if v == targetSum{
	//		return true
	//	}
	//}
	//return false
}
//func getAllPathSum(root *TreeNode)[]int{
//	if root == nil {
//		return nil
//	}
//	// 树不为空
//	var result  =  make([]int,0)
//	left := getAllPathSum(root.Left)
//	right := getAllPathSum(root.Right)
//	if left == nil || right == nil{
//		// 判断该树中是否存在 根节点到叶子节点 的路径
//		if left == nil && right == nil{
//			result = append(result, root.Val)
//		}
//		if left !=nil {
//			for _,v :=range left{
//				result = append(result,v + root.Val)
//			}
//		}
//		if right !=nil {
//			for _,v :=range right{
//				result = append(result,v + root.Val)
//			}
//		}
//	}else{
//		for _,v :=range left{
//			result = append(result,v + root.Val)
//		}
//		for _,v :=range right{
//			result = append(result,v + root.Val)
//		}
//	}
//	return result
//}
//leetcode submit region end(Prohibit modification and deletion)
