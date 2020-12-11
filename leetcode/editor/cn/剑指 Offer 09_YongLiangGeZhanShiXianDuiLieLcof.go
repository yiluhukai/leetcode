//用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的
//功能。(若队列中没有元素，deleteHead 操作返回 -1 )
//
//
//
// 示例 1：
//
// 输入：
//["CQueue","appendTail","deleteHead","deleteHead"]
//[[],[3],[],[]]
//输出：[null,null,3,-1]
//
//
// 示例 2：
//
// 输入：
//["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
//[[],[],[5],[2],[],[]]
//输出：[null,-1,null,null,5,2]
//
//
// 提示：
//
//
// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用
//
// Related Topics 栈 设计
// 👍 146 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	data []int
	size int
	top  int // 代表下次插入元素的位置，  top ==  tail  || size ==0  是空栈
	tail int
}

func (st *MyStack) Push(elem int) {
	if st.size == len(st.data) {
		st.data = append(st.data, elem)
	} else {
		st.data[st.top] = elem
	}
	st.size++
	st.top++
}

func (st *MyStack) Pop() int {
	elem := st.Peek()
	st.top--
	st.size--
	return elem
}

//  获取栈顶元素
func (st *MyStack) Peek() int {
	if st.top == st.tail {
		panic("error")
	}
	elem := st.data[st.top-1]
	return elem
}

//  判断一个栈是不是空栈
func (st *MyStack) IsEmpty() bool {
	if st.size == 0 {
		return true
	} else {
		return false
	}
}


type CQueue struct {
	pushStack MyStack
	popStack  MyStack
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.pushStack.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.popStack.IsEmpty(){
		for !this.pushStack.IsEmpty(){
			x:=this.pushStack.Pop()
			this.popStack.Push(x)
		}
	}
	if this.popStack.IsEmpty(){
		return -1
	}
	return this.popStack.Pop()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
//leetcode submit region end(Prohibit modification and deletion)
