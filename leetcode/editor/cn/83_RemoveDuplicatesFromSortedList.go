//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。 
//
// 示例 1: 
//
// 输入: 1->1->2
//输出: 1->2
// 
//
// 示例 2: 
//
// 输入: 1->1->2->3->3
//输出: 1->2->3 
// Related Topics 链表 
// 👍 438 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
	解法：遍历链表，将链表中相同的元素删除
 */
// 解答成功:
// 执行耗时:4 ms,击败了86.50% 的Go用户
// 内存消耗:3.1 MB,击败了62.03% 的Go用户

func deleteDuplicates(head *ListNode) *ListNode {
	for cur:=head;cur!=nil && cur.Next!=nil;{
		if cur.Val == cur.Next.Val{
			cur.Next =  cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return head
}
//leetcode submit region end(Prohibit modification and deletion)


