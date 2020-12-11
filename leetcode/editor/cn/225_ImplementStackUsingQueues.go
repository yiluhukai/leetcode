//使用队列实现栈的下列操作： 
//
// 
// push(x) -- 元素 x 入栈 
// pop() -- 移除栈顶元素 
// top() -- 获取栈顶元素 
// empty() -- 返回栈是否为空 
// 
//
// 注意: 
//
// 
// 你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合
//法的。 
// 你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。 
// 你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。 
// 
// Related Topics 栈 设计 
// 👍 245 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)

//思路：用两个队列来模拟栈，队列1用来存放数据，队列2 用来辅助队列1实现每次新增的数据都在队列的头部。 这样子就可以实现栈的后进先出。
//
//入栈：现将元素放入队列2中。然后将队列1中的元素的出队然后放入队列2中（为空则不进行），然后将队列2中的元素出队列放入队列1中。这样子都报证每次
//插入栈中的元素都在队列1的队首位置。

//  golang中没有队列，所以我们可以使用切片来模拟


//type MyStack struct {
//	Queue1 []int
//	Queue2 []int
//}
//
//
///** Initialize your data structure here. */
//func Constructor() MyStack {
//	return MyStack{}
//}
//
//
///** Push element x onto stack. */
//func (this *MyStack) Push(x int)  {
//	this.Queue2 =  append(this.Queue2,x)
//	if len(this.Queue1)!=0 {
//		this.Queue2  = append(this.Queue2,this.Queue1...)
//	}
//	this.Queue1 =  this.Queue2
//	this.Queue2 =  make([]int,0)
//}
//
//
///** Removes the element on top of the stack and returns that element. */
//func (this *MyStack) Pop() int {
//	x :=  this.Queue1[0]
//	this.Queue1 = this.Queue1[1:]
//	return x
//}
//
//
///** Get the top element. */
//func (this *MyStack) Top() int {
//	return this.Queue1[0]
//}
//
//
///** Returns whether the stack is empty. */
//func (this *MyStack) Empty() bool {
//	return len(this.Queue1) == 0
//}

/*
	思路:用队列实现栈的主要思想在于让进入队列的元素能像栈一样，后进先出，为此需要让最后进栈的元素放到队列的最前面。
	我们使用一个队列来实现时，当一个元素进入队列时，我们需要让当前队列的所有元素都出队列然后在进队列。
*/



type MyStack struct {
	Queue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	size:= len(this.Queue)
	this.Queue =  append(this.Queue,x)
	for i:=size;i>0;i--{
		n:=this.Pop()
		this.Queue = append(this.Queue,n)
	}
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x :=  this.Queue[0]
	this.Queue = this.Queue[1:]
	return x
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Queue[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)


