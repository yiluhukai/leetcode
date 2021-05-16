package main

import "container/list"

//给定一个 N 叉树，找到其最大深度。
//
// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
//
// N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,null,3,2,4,null,5,6]
//输出：3
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,
//null,13,null,null,14]
//输出：5
//
//
//
//
// 提示：
//
//
// 树的深度不会超过 1000 。
// 树的节点数目位于 [0, 104] 之间。
//
// Related Topics 树 深度优先搜索 广度优先搜索
// 👍 162 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:3.7 MB,击败了5.09% 的Go用户
func maxDepth(root *Node) int {
	//  BFS
	if root == nil {
		return 0
	}
	if root.Children == nil || len(root.Children) == 0 {
		return 1
	}
	//非空N叉树
	l := list.New()
	l.PushBack(root)
	depth := 0
	for length := l.Len(); length != 0; length = l.Len() {
		// 处理每层的节点
		for i := 0; i < length; i++ {
			e := l.Remove(l.Front()).(*Node)
			if e.Children != nil && len(e.Children) != 0 {
				for _, kid := range e.Children {
					l.PushBack(kid)
				}
			}
		}
		depth++
	}
	return depth
}

//执行耗时:0 ms,击败了100.00% 的Go用户
//内存消耗:3.5 MB,击败了7.17% 的Go用户
//func maxDepth(root *Node) int {
//    //  DFS
//	if root == nil {
//		return 0
//	}
//	if root.Children == nil || len(root.Children) == 0{
//		return 1
//	}
//	//非空N叉树
//	res := []int{}
//	for _,kid :=range root.Children{
//		res = append(res, maxDepth(kid))
//	}
//	return 1 + max(res)
//}

func max(children []int) int {
	max := children[0]
	for _, kid := range children {
		if kid > max {
			max = kid
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
