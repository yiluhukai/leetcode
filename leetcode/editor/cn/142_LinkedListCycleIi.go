//给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，po
//s 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
//
// 说明：不允许修改给定的链表。
//
// 进阶：
//
//
// 你是否可以使用 O(1) 空间解决此题？
//
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [3,2,0,-4], pos = 1
//输出：返回索引为 1 的链表节点
//解释：链表中有一个环，其尾部连接到第二个节点。
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2], pos = 0
//输出：返回索引为 0 的链表节点
//解释：链表中有一个环，其尾部连接到第一个节点。
//
//
// 示例 3：
//
//
//
//
//输入：head = [1], pos = -1
//输出：返回 null
//解释：链表中没有环。
//
//
//
//
// 提示：
//
//
// 链表中节点的数目范围在范围 [0, 104] 内
// -105 <= Node.val <= 105
// pos 的值为 -1 或者链表中的一个有效索引
//
// Related Topics 链表 双指针
// 👍 790 👎 0
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
	暴力解法：将链表中的节点放入到队列中，遍历队列中是否包含这个节点的引用,如果有环，则能在队列中找到这个节点，
	且当前节点和队列中的节点匹配时，证明有环，且当前节点就是环的入口。
*/
//解答成功:
//执行耗时:44 ms,击败了12.72% 的Go用户
//内存消耗:4.7 MB,击败了24.07% 的Go用户
//func detectCycle(head *ListNode) *ListNode {
//	sli := make([]*ListNode, 0)
//	cur := head
//	for cur != nil {
//		//  检测sli没有和cur相同的节点
//		for _, val := range sli {
//			if val == cur {
//				return cur
//			}
//		}
//		sli = append(sli, cur)
//		cur = cur.Next
//	}
//	return nil
//}
/*
	优化解法：使用快慢指针去匹配链表是否有环，当有环的同时,
	当有环的时候，slow == fast 且这个点就是相遇点meet
	head - entry = s , entry - meet = f,
	slow 走过的路程为L,则fast 走过的路程为2L,环的大小为r
	L =  s +f ,2L = s+nr+f-1,起点对一个。
	-> s+f = nr-1 , 可以看出，一个指针从head开始，另一个指针从slow的下一个位置开始，经过s+f步会在slow的位置再次相遇。
	所以也会在起点的位置相遇。
*/

//解答成功:
//执行耗时:4 ms,击败了97.74% 的Go用户
//内存消耗:3.7 MB,击败了100.00% 的Go用户
func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next

	//只判断fast 指针是否指向nil

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 相遇的点，证明有环
			break
		}
	}
	//  有环,不能用fast !=nil 判断，当两边质量两个切没有环的时候，fast!=nil
	if fast == slow {
		cur := head
		slow = slow.Next
		for cur != slow {
			cur = cur.Next
			slow = slow.Next
		}
		return cur
	}
	//  没有环
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
