//给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。 
//
// 示例 1: 
//
// 输入: 1->2->3->3->4->4->5
//输出: 1->2->5
// 
//
// 示例 2: 
//
// 输入: 1->1->1->2->3
//输出: 2->3 
// Related Topics 链表 
// 👍 412 👎 0

/*

 */
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
	思想：首先添加一个头节点，从head节点开始比较，pre指向数据节点的前一个节点
	当两个节点的数据相等时，标志为设置为true,head = head.Next
	当两个节点数据不同时，
		如果标志位为true.这时需要 pre指向head.Next节点，head.Next=nil, head= pre.Next,标志为设置为false
		如果标志为为false,更新pre和head
	还有一种特殊情况，就是链表最后几个节点的等于一个数，这时的标志位为true,但是我们没有去处理，需要在循环外去处理
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head ==nil || head.Next ==nil {
		return head
	}
	// 起码有两个节点，创建一个头节点
	h := new(ListNode)
	h.Next = head
	//  记录前一个节点的指针
	pre:= h
	isRepeat := false

	for head!=nil && head.Next!=nil{
		if head.Val == head.Next.Val {
			isRepeat = true
			head =head.Next
		}else{
			if isRepeat {
				pre.Next = head.Next
				head.Next = nil
				head = pre.Next
				isRepeat = false

			}else{
				pre = head
				head =head.Next
			}
		}

	}
	if isRepeat && head !=nil{
		pre.Next = head.Next
		head.Next = nil
		head = pre.Next
		//isRepeat = false
	}
	next:= h.Next
	// 取出头节点
	h.Next = nil
	return next
}
//leetcode submit region end(Prohibit modification and deletion)


