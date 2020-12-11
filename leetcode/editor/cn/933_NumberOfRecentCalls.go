//写一个 RecentCounter 类来计算特定时间范围内最近的请求。
//
// 请你实现 RecentCounter 类：
//
//
// RecentCounter() 初始化计数器，请求数为 0 。
// int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求
//）。确切地说，返回在 [t-3000, t] 内发生的请求数。
//
//
// 保证 每次对 ping 的调用都使用比之前更大的 t 值。
//
//
//
// 示例：
//
//
//输入：
//["RecentCounter", "ping", "ping", "ping", "ping"]
//[[], [1], [100], [3001], [3002]]
//输出：
//[null, 1, 2, 3, 3]
//
//解释：
//RecentCounter recentCounter = new RecentCounter();
//recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
//recentCounter.ping(100);   // requests = [1, 100]，范围是 [-2900,100]，返回 2
//recentCounter.ping(3001);  // requests = [1, 100, 3001]，范围是 [1,3001]，返回 3
//recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002]，范围是 [2,3002]，返回
//3
//
//
//
//
// 提示：
//
//
// 1 <= t <= 109
// 保证每次对 ping 调用所使用的 t 值都 严格递增
// 至多调用 ping 方法 104 次
//
// Related Topics 队列
// 👍 70 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)

//const max = 3002

//var (
//	arr =  [max]int{0}
//	start =0
//	end =0
//)

//type RecentCounter struct {
//	 sli []int
//}

//type RecentCounter struct {
//	//sli []int
//	arr   [max]int
//	start int
//	end   int
//}


type Queue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	data int
	next  *Node
}
// 添加元素
func(q * Queue)add(x int){
	newNode := Node{
		data:x,
	}
	// 空队列
	if q.tail == nil{
		q.tail = &newNode
		q.head = &newNode
	}else{
		//  队列不为空
		last := q.tail
		last.next =  &newNode
		q.tail = &newNode
	}
	q.size++
}
// 删除元素
func(q * Queue)poll(){
	//空队列
	if q.head == nil{
		panic("Queue is empty")
	}
	head := q.head
	//  只有一个元素
	q.head = head.next
	if head.next == nil {
		q.tail =nil
	}else{
		head.next =nil
	}

	q.size--
}

func (q *Queue)getHeadVal()int{
	return q.head.data
}

type RecentCounter struct {
   Q Queue
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	/*暴力解法：将所有调用ping()的时刻存起来,最多需要10^4的空间，然后每次反向遍历数组去统计符合条件的个数*/
	//解答成功:
	//执行耗时:160 ms,击败了53.13% 的Go用户
	//内存消耗:8.9 MB,击败了64.06% 的Go用户
	//	this.sli = append(this.sli, t)
	//	count:=0
	//	length:= len(this.sli)
	//	//  当数组的所有元素符合条件时，end会小于0
	//	for end:=length-1;end>=0&&this.sli[end]>=t-3000;end--{
	//		count++
	//	}
	//	//fmt.Printf("t=%v,count=%v\n",t,count)
	//	return count

	/*
		优化解法：使用双指针方法：开始指针指向的是符合条件的下标，结束指针指向的是下次插入元素的位置，
		最多的时候需要保存3001个元素 [0,3000],所以需要的存储空间为3002，负责会覆盖end指针会覆盖start指针
	*/
	//this.arr[this.end] = t
	//this.end = (this.end+1) % max
	//
	//for this.arr[this.start] < t-3000 {
	//	this.start =  (this.start+1)%max
	//}
	//if this.end >this.start {
	//	return this.end - this.start
	//}else{
	//	return max - (this.start-this.end)
	//}

	/*
		最优解法：使用队列替换上面的数组，这样子可以有效的减小空间复杂度
		队列中存放的是符合条件的请求
	*/
	// 放入队列中
	this.Q.add(t)
	//  从队列中删除小于 t-3000的元素
	for this.Q.getHeadVal()  < t-3000 {
		//  出队操作
		this.Q.poll()
	}
	return  this.Q.size
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
//leetcode submit region end(Prohibit modification and deletion)

//func main(){
//	rc:= Constructor()
//	testCase:=[]int{1178,1483,1563,4054,4152}
//	for _,val :=range testCase{
//		res:=rc.Ping(val)
//		fmt.Printf("res=%v\n",res)
//	}
//}
