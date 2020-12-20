//不使用任何库函数，设计一个跳表。 
//
// 跳表是在 O(log(n)) 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，并且跳表的代码长度相较下更短，其设计思想
//与链表相似。 
//
// 例如，一个跳表包含 [30, 40, 50, 60, 70, 90]，然后增加 80、45 到跳表中，以下图的方式操作： 
//
// 
//Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons 
//
// 跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。跳表的每一个操作的平均时间复杂度是 O(log(
//n))，空间复杂度是 O(n)。 
//
// 在本题中，你的设计应该要包含这些函数： 
//
// 
// bool search(int target) : 返回target是否存在于跳表中。 
// void add(int num): 插入一个元素到跳表。 
// bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。 
//
// 
//
// 了解更多 : https://en.wikipedia.org/wiki/Skip_list 
//
// 注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。 
//
// 样例: 
//
// Skiplist skiplist = new Skiplist();
//
//skiplist.add(1);
//skiplist.add(2);
//skiplist.add(3);
//skiplist.search(0);   // 返回 false
//skiplist.add(4);
//skiplist.search(1);   // 返回 true
//skiplist.erase(0);    // 返回 false，0 不在跳表中
//skiplist.erase(1);    // 返回 true
//skiplist.search(1);   // 返回 false，1 已被擦除
// 
//
// 约束条件: 
//
// 
// 0 <= num, target <= 20000 
// 最多调用 50000 次 search, add, 以及 erase操作。 
// 
// Related Topics 设计 
// 👍 46 👎 0
package main

import (
	"fmt"
	"math"
	"math/rand"
)

//leetcode submit region begin(Prohibit modification and deletion)
//	链表的节点
type Node_t struct {
	Right *Node_t
	Down  *Node_t
	Data int
}
type Skiplist struct {
	Head *Node_t  //  指向跳表的开始位置的
	Level int // 跳表的层级
	Length int //  跳表的原链表的长度
}



func Constructor() Skiplist {
	return Skiplist{
		Level: 1,
		// 初始化头节点
		Head: &Node_t{Data: -1},
	}
}

 // 从跳表的头节点开始，逐层查找
func (this *Skiplist) Search(target int) bool {
	cur :=  this.Head
	for cur !=nil {
		//  当同层的右侧节点不非空的时候，继续向右侧查找
		for cur.Right !=nil && cur.Right.Data < target {
			cur = cur.Right
		}
		if cur.Right !=nil && cur.Right.Data  == target {
			return true
		}
		// 向下层查找
		cur =  cur.Down
	}
	// 跳表的每一层都找了，没有 == target的节点
	return false
}
/*
	插入一个元素到跳表。
	思路：找到要插入元素的位置(链表是有序的，原链表上第一个不大于这个数的前面)，然后采用抛硬币的方式，判断是否生成索引
*/
//
//解答成功: 无索引的情况
//执行耗时:668 ms,击败了18.60% 的Go用户
//内存消耗:9.7 MB,击败了76.74% 的Go用户

func (this *Skiplist) Add(num int)  {
	cur :=  this.Head
	//  存储每层的cur
	var curs = make([]*Node_t,0)
	// 使用切片存储当前曾的cur,cur指向的是要建立索引的元素
	for cur !=nil {
		//  当同层的右侧节点不非空的时候，继续向右侧查找
		for cur.Right !=nil && cur.Right.Data < num {
			cur = cur.Right
		}
		curs = append(curs,cur)
		// 向下层查找
		cur =  cur.Down
	}
	// 添加元素到原链表上
	size :=len(curs)
	if size >0 {
		newNode := &Node_t{
			Data: num,
		}
		last := size-1
		cur = curs[last]
		newNode.Right = cur.Right
		cur.Right = newNode
		this.Length++
		//  添加索引
		this.addIndexForSkiplist(newNode,curs,last)
	}
}
// 为跳表的元素添加索引
func(this * Skiplist) addIndexForSkiplist(target *Node_t, sli []*Node_t,size int){
	//  先抛硬币确定是否添加索引
	node := target
	coins := rand.Intn(2)
	//  限制跳表的层级
	for coins == 1 && this.Level <=  int(math.Log2(float64(this.Length))) {
		if size > 0{
			// 在原来索引的基础上添加索引
			newNode := &Node_t{
				Data: node.Data,
				Down: node,
			}
			size--
			newNode.Right = sli[size].Right
			sli[size].Right =  newNode
			// 更新node
			node = newNode
			coins = rand.Intn(2)
		}else{
			//  新建新的索引
				newNode := &Node_t{
					Data: node.Data,
					Down: node,
				}
				head := &Node_t{
					Data: -1,
					Down: this.Head,
					Right: newNode,
				}
				this.Head = head
				this.Level++
		}
	}

}

/*
	在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。
	思路：找到前一个要删除的前一个节点，移除当前节点以及下层的节点
*/

func (this *Skiplist) Erase(num int) bool {
	cur :=  this.Head
	flag := false
	for cur !=nil {
		//  当同层的右侧节点不非空的时候，继续向右侧查找
		for cur.Right !=nil && cur.Right.Data < num {
			cur = cur.Right
		}
		if cur.Right !=nil && cur.Right.Data  == num {
			// cur 为要删除的节点前一个节点
			right :=  cur.Right
			cur.Right = right.Right
			right.Right = nil
			right.Down = nil
			flag =true
		}
		// 向下层查找
		cur =  cur.Down
	}
	if flag == true {
		this.Length --
	}
	return flag
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
//leetcode submit region end(Prohibit modification and deletion)
