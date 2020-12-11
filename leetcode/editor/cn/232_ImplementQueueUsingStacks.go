//请你仅使用两个栈实现先入先出队列。队列应当支持一般队列的支持的所有操作（push、pop、peek、empty）： 
//
// 实现 MyQueue 类： 
//
// 
// void push(int x) 将元素 x 推到队列的末尾 
// int pop() 从队列的开头移除并返回元素 
// int peek() 返回队列开头的元素 
// boolean empty() 如果队列为空，返回 true ；否则，返回 false 
// 
//
// 
//
// 说明： 
//
// 
// 你只能使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
// 
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。 
// 
//
// 
//
// 进阶： 
//
// 
// 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。 
// 
//
// 
//
// 示例： 
//
// 
//输入：
//["MyQueue", "push", "push", "peek", "pop", "empty"]
//[[], [1], [2], [], [], []]
//输出：
//[null, null, null, 1, 1, false]
//
//解释：
//MyQueue myQueue = new MyQueue();
//myQueue.push(1); // queue is: [1]
//myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
//myQueue.peek(); // return 1
//myQueue.pop(); // return 1, queue is [2]
//myQueue.empty(); // return false
// 
//
// 
// 
//
// 
//
// 提示： 
//
// 
// 1 <= x <= 9 
// 最多调用 100 次 push、pop、peek 和 empty 
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作） 
// 
// Related Topics 栈 设计 
// 👍 255 👎 0
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

type MyQueue struct {
	pushStack MyStack
	popStack  MyStack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.pushStack.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.popStack.IsEmpty(){
		for !this.pushStack.IsEmpty(){
			x:=this.pushStack.Pop()
			this.popStack.Push(x)
		}
	}
	if this.popStack.IsEmpty(){
		panic("queue is Empty")
	}
	return this.popStack.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.popStack.IsEmpty(){
		for !this.pushStack.IsEmpty(){
			x:=this.pushStack.Pop()
			this.popStack.Push(x)
		}
	}
	if this.popStack.IsEmpty(){
		panic("queue is Empty")
	}
	return this.popStack.Peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.popStack.IsEmpty() && this.pushStack.IsEmpty()
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)


