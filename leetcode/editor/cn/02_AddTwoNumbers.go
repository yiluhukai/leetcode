//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
// Related Topics 链表 数学
// 👍 5308 👎 0
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l * ListNode)addNode(val int){
	newNode := new(ListNode)
	newNode.Val =  val
	l.Next =newNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 */


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1,p2 := l1 ,l2
	res := new(ListNode)
	cur := res
	// 进位
	high := 0

	for p1!=nil || p2!=nil {
		num1,num2:= 0,0
		if p1 != nil{
			num1 = p1.Val
		}
		if p2 !=nil{
			num2 = p2.Val
		}
		sum:=num1 +num2 + high

		low := sum%10
		high =  sum /10



		//  创建一个新的节点
		//cur.addNode(low)

		newNode := new(ListNode)
		newNode.Val =  low
		cur.Next =newNode
		// 指针向后走
		if p1!=nil {
			p1 =p1.Next
		}
		if p2 !=nil{
			p2 =p2.Next
		}
		cur =cur.Next
	}
	// 当两个链表遍历结束，判断有进位
	if high != 0 {
		//cur.addNode(high)
		newNode := new(ListNode)
		newNode.Val =  high
		cur.Next =newNode
	}
	return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	testArr1 := [...]int{2,4,9}
//	testArr2 := [...]int{5,6,4}
//
//	l1,l2 := new(ListNode),new(ListNode)
//
//	p1,p2 :=l1,l2
//
//	for i:=0;i<len(testArr1);i++ {
//		p1.addNode(testArr1[i])
//		p2.addNode(testArr2[i])
//		p1 =p1.Next
//		p2 =p2.Next
//	}
//
//	for res:=addTwoNumbers(l1.Next,l2.Next);res!=nil;res =res.Next{
//		fmt.Printf("%v\t",res.Val)
//	}
//}
