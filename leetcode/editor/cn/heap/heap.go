package heap

type  Heap struct {
	array []int
}


/**
	构建小顶堆：输入是一颗完全二叉树 -> 转换成一个小顶堆
    思路：从第一个非叶子节点开始调整，让所有非叶子节点下沉
*/

func BuildHeap(array []int) *Heap {
	length := len(array)

	for i:=length/2-1;i>=0;i--{
		downAdjust(array,i,length)
	}
	return &Heap{array: array}
}


/**
插入一个节点: 如何当前节点大于父节点，节点上浮；
所以时间复杂度：N是节点的个数 logN代表二叉堆的高度  O(logN)
*/
func(h  *Heap)Insert(val int){
	h.array = append(h.array, val)
	upAdjust(h.array)
}


/**
删除一个堆顶元素: 删除堆顶元素，然后使用堆中的最后一个元素作为堆顶元素，接下来自上而下调整堆
N是节点的个数 logN代表了二叉堆的高度 O(logN)
*/
func (h *Heap)Delete(){
	length:=len(h.array)
	if length < 1{
		return
	}

	// 删除第一个节点，用最后一个节点作为头节点
	h.array[0] = h.array[length-1]
	// 删除最后一个节点
	h.array = h.array[:length-1]

	length = len(h.array)
	if length < 2{
		return
	}
	downAdjust(h.array,0,length)
}

/**
节点的上浮操作:从最后一个节点开始，如果当前的父节点小于当前节点，则交换两个节点的值，知道当前节点指向根结点
*/
func upAdjust(array []int){
	length :=  len(array)
	if length < 1 {
		return
	}
	childIndex := length-1
	parentIndex := (childIndex-1) /2
	temp := array[childIndex]
	for childIndex > 0 &&  array[parentIndex] > temp{
		array[childIndex] =  array[parentIndex]
		childIndex = parentIndex
		parentIndex =  (childIndex - 1)/2
	}
	// 当到达根结点或者父节点不大于子节点时
	array[childIndex] = temp
}

/**
节点的下沉操作
*/

func downAdjust(array []int,parentIndex,length int){
	// 比较当前节点和子节点的值,如果父节点大于子节点中较小的值，进行交换
	childIndex := parentIndex *2 + 1
	temp := array[parentIndex]
	for childIndex  < length{
		// 右子节点存在切小于左子节点
		if childIndex + 1 <length && array[childIndex + 1] < array[childIndex]{
			childIndex ++
		}
		// 不用下沉
		if temp <= array[childIndex]{
			break
		}
		array[parentIndex] =  array[childIndex]
		parentIndex = childIndex
		childIndex = childIndex * 2+1
	}
	array[parentIndex] =  temp
}


