package  main

import "container/list"

//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//
// 
//
// 示例： 
//二叉树：[3,9,20,null,null,15,7], 
//
// 
//    3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回其层序遍历结果： 
//
// 
//[
//  [3],
//  [9,20],
//  [15,7]
//]
// 
// Related Topics 树 广度优先搜索 
// 👍 844 👎 0


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
//内存消耗:2.9 MB,击败了25.63% 的Go用户
func levelOrder(root *TreeNode) [][]int {
	// 判断root是否为空
	if root == nil{
		return nil
	}
	// 不为空的情况：创建一个切片保存每一层的切片
    l := list.New()
    l.PushBack(root)
    res := make([][]int,0)
    for length:= l.Len(); length!=0;length= l.Len(){
    	// 对当前层的元素作遍历
    	cur := []int{}
    	for i:=0;i<length;i++{
    		e:=l.Remove(l.Front()).(*TreeNode)
			cur = append(cur,e.Val)
			//加入当前节点的左右节点
			if e.Left !=nil{
				l.PushBack(e.Left)
			}
			if e.Right !=nil{
				l.PushBack(e.Right)
			}
		}
		// 将该层的节点加入到res切片中
		res = append(res, cur)
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
