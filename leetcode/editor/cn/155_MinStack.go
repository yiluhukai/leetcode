//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
//
//
//
//
// 示例:
//
// 输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// pop、top 和 getMin 操作总是在 非空栈 上调用。
//
// Related Topics 栈 设计
// 👍 737 👎 0

package main

////leetcode submit region begin(Prohibit modification and deletion)
//type MinStack struct {
//	data     []int
//	top      int
//	//tail     int
//	minIndex int
//	size  int
//}
//
///** initialize your data structure here. */
////func Constructor() MinStack {
////	return MinStack{
////
////	}
////}
//
//func (this *MinStack) Push(x int) {
//	if this.size == len(this.data){
//		this.data= append(this.data,x)
//	}else{
//		this.data[this.top] = x
//	}
//	//  非空栈
//	if x < this.data[this.minIndex] {
//		this.minIndex = this.top
//	}
//	this.size++
//	this.top++
//}
//
//func (this *MinStack) Pop() {
//	this.top--
//	this.size--
//	//  最小的是栈顶元素
//	if this.minIndex == this.top {
//		min := this.data[0]
//		this.minIndex = 0
//		for index,length:=0,this.size;index < length;index++ {
//			if min > this.data[index] {
//				this.minIndex = index
//				min =  this.data[index]
//			}
//		}
//	}
//
//}
//
//func (this *MinStack) Top() int {
//	return this.data[this.top-1]
//}
//
//func (this *MinStack) GetMin() int {
//	return this.data[this.minIndex]
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)

//func main()  {
//	st :=  Constructor()
//	st.Push(2147483646)
//	st.Push(2147483646)
//	st.Push(2147483647)
//
//	fmt.Printf("top=%v\n",st.Top())
//	st.Pop()
//	fmt.Printf("min=%v\n",st.GetMin())
//	st.Pop()
//	fmt.Printf("min=%v\n",st.GetMin())
//	st.Pop()
//	st.Push(2147483647)
//	fmt.Printf("top=%v\n",st.Top())
//	fmt.Printf("min=%v\n",st.GetMin())
//	st.Push(-2147483648)
//	fmt.Printf("top=%v\n",st.Top())
//	fmt.Printf("min=%v\n",st.GetMin())
//	st.Pop()
//	fmt.Printf("min=%v\n",st.GetMin())
//}


//思路：使用两个队列实现最小栈
//class MinStack {
//	Deque<Integer> xStack; //主栈，存放所有入栈数据
//	Deque<Integer> minStack; //每次入栈当前最小数据
//	public MinStack() {
//	xStack = new LinkedList<Integer>();
//	minStack = new LinkedList<Integer>();
//	minStack.push(Integer.MAX_VALUE);//默认最小数据为整形最大值
//	}
//
//	public void push(int x) {
//	xStack.push(x);//入栈
//	minStack.push( minStack.peek()< x? minStack.peek():x); //将当前最小值入栈，为保证数量一致，所以会重复存放最小值
//	}
//
//	public void pop() {
//	xStack.pop();//出栈
//	minStack.pop();//当前最小值也出栈
//	}
//
//	public int top() {
//	return xStack.peek();//获取栈顶元素
//	}
//
//	public int getMin() {
//	return minStack.peek();//通过最小栈获取当前最小元素
//	}
//}
