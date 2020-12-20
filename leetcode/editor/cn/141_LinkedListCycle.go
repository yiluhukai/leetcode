//给定一个链表，判断链表中是否有环。 
//
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的
//位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。 
//
// 如果链表中存在环，则返回 true 。 否则，返回 false 。 
//
// 
//
// 进阶： 
//
// 你能用 O(1)（即，常量）内存解决此问题吗？ 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
// 
//
// 示例 2： 
//
// 
//
// 输入：head = [1,2], pos = 0
//输出：true
//解释：链表中有一个环，其尾部连接到第一个节点。
// 
//
// 示例 3： 
//
// 
//
// 输入：head = [1], pos = -1
//输出：false
//解释：链表中没有环。
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点的数目范围是 [0, 104] 
// -105 <= Node.val <= 105 
// pos 为 -1 或者链表中的一个 有效索引 。 
// 
// Related Topics 链表 双指针 
// 👍 875 👎 0

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
	解题思路：设置两个快慢指针，快指针每次走两部，慢指针每次走一步。如果有环，那么这两个指针一定会相遇。且是在慢指针进入环的第一圈相遇。
	为啥会相遇且是在慢指针进入环的第一圈：假如有环，那么快指针肯定在环上，当慢指针进入环后，他们的最大距离是快指针在慢指针的前面，最小距离就是他们
	都在入口位置。假设环的大小是n,最大距离是n-1，由于快指针每次比慢指针多走一步，随意每走一步，他们的距离减小一步，最对n-1步他们就会相遇。
 */

//解答成功:
//执行耗时:4 ms,击败了98.16% 的Go用户
//内存消耗:3.6 MB,击败了65.21% 的Go用户
//func hasCycle(head *ListNode) bool {
//    slow := head
//    fast:=head
//    if head == nil{
//    	return false
//	}else{
//		fast = head.Next
//	}
//	//只判断fast 指针是否指向nil
//
//	for fast != nil && fast.Next!=nil && fast.Next.Next != nil{
//		fast = fast.Next.Next
//		slow = slow.Next
//		if fast  == slow {
//			return true
//		}
//	}
//
//	return false
//}

/*暴力解法: 用一个队列去保存链表的节点，当节点在队列中再次出现的时候，代表链表有环*/
//
//解答成功:
//执行耗时:48 ms,击败了11.21% 的Go用户
//内存消耗:4.7 MB,击败了17.36% 的Go用户
func hasCycle(head *ListNode) bool {
	sli := make([]*ListNode,0)
	cur :=head
	for cur !=nil{
		//  检测sli没有和cur相同的节点
		for _,val :=range sli{
			if val == cur{
				return true
			}
		}
		sli  =append(sli,cur)
		cur = cur.Next
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)


