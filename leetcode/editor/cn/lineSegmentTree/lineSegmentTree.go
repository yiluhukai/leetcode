package lineSegmentTree

type LineSegmentTree struct {
	// 存储线段树中元素的节点
	container []int
	n  int// 区间长度
	size int // 二叉树的节点个数
	nums []int // 存储原始点的数组
}
// 假设安全二叉树有x个节点，那么他的叶子节点个数为
// 最后一个节点的父节点的下标为  （x-1）-((x-1-1)/2)) =  n ,那么x === 2n
// 验证一下：n == 3
func NewLineSegmentTree(n int, nums []int) *LineSegmentTree{
	l := &LineSegmentTree{
		nums:nums,
		n:n,
		// 完全二叉树的叶子节点的个数等n,所以二叉树的节点个数为
		container: make([]int,n * 2),
	}
	l.build(0,0,n-1)
	l.container = l.container[:l.size]
	return l
}

//非叶子节点[a,b]，它的左儿子表示的区间为[a,(a+b)/2]，右儿子表示的区间为[(a+b)/2+1,b]
// index代表节点在线段树中的下标，left 和 right代表节点的左右区间值
func (l *LineSegmentTree)build(index,left,right int){
	l.size++
	if left == right{
		l.container[index] = l.nums[left]
		return
	}
	mid := (left + right) /2
	leftChildIndex,rightChildIndex := 2 * index +1,2*index+2
	l.build( leftChildIndex,left,mid)
	l.build( rightChildIndex,mid+1,right)
	l.container[index] = l.container[leftChildIndex] + l.container[rightChildIndex]
}


// 对区间求和[qLeft,qRight]的和
// index代表线段树中节点的下标，left和right代表节点的区间值[left,right]
func (l *LineSegmentTree)Query(index,left,right,qLeft,qRight int)int{
	// 两个区间没有交集
	if qLeft > right || qRight < left {
		return 0
	}
	// 求和区间包含当前的区间
	if qRight >= right && qLeft <= left{
		return l.container[index]
	}
	// 【left,right】和【qLeft，qRight】两个区间有交集，将[left,right]拆分成[left,mid] [mid+1,right]再去求和
	mid := (right +left) / 2
	leftSum := l.Query(index *2 +1,left,mid,qLeft,qRight)
	rightSum := l.Query(index*2 +2 ,mid+1,right,qLeft,qRight)
	return leftSum + rightSum
}

// 修改nums中的值，对线段树对应的节点做更新
// index代表线段树中节点的位置,节点的区间值[left,right]，
// position代表更新了的下标，value代表更新后的值
func (l *LineSegmentTree)Update(index,left,right,position,value int){
	// 如果是叶子结点
	if left == right && position == left {
		l.container[index] = value
		return
	}
	// 非叶子节点，递归去更新子节点的值，然后重新计算当前节点的值
	mid := (left + right)/2
	if position <= mid{
		//更新左子树
		l.Update(index * 2 + 1,left,mid,position,value)
	}else{
		l.Update(index * 2 +2,mid +1,right,position,value)
	}
	// 更新当前节点的和
	l.container[index] = l.container[index * 2 + 1] + l.container[index * 2 + 2]
}

