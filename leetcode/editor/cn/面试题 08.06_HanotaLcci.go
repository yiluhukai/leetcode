//在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只
//能放在更大的盘子上面)。移动圆盘时受到以下限制: 
//(1) 每次只能移动一个盘子; 
//(2) 盘子只能从柱子顶端滑出移到下一根柱子; 
//(3) 盘子只能叠在比它大的盘子上。 
//
// 请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。 
//
// 你需要原地修改栈。 
//
// 示例1: 
//
//  输入：A = [2, 1, 0], B = [], C = []
// 输出：C = [2, 1, 0]
// 
//
// 示例2: 
//
//  输入：A = [1, 0], B = [], C = []
// 输出：C = [1, 0]
// 
//
// 提示: 
//
// 
// A中盘子的数目不大于14个。 
// 
// Related Topics 递归 
// 👍 69 👎 0
package main

/*
	使用循环解决汉诺塔问题：
	按照美国学者的思路：汉诺塔问题可以分两部去完成，在进行这两步之前，需要一些准备工作。
	将三根柱子摆成品字形，A在最上面，当盘子数量为奇数时，柱子的顺序顺时针为ACB. ->1
	当盘子的个数为偶数时,盘子的顺序顺时针为ABC.

	接下来重复两个步骤：
	1. 将最小的盘子移动到一个柱子上(顺时针的方向)
	2. 比较另外两个柱子上，当有一个柱子为空时，将盘子从非空的移动到空的柱子上，
	当两个都非空时，把较小的那个移动较大哪一个柱子上。
	重复上面两个动作，直到所以的盘子都移动到C柱子上。
 */
//解答成功:
//执行耗时:0 ms,击败了100.00% 的Go用户
//  内存消耗:1.9 MB,击败了16.48% 的Go用户
//leetcode submit region begin(Prohibit modification and deletion)
/*
	使用递归的思想去解决这个问题：
	把N个盘子中的N-1个盘子看成一个整体，先把N-1个盘子移动到中间柱，再把第N个盘子移动到目标柱。
	接下来再把N-1个盘子移动移动到目标柱。
	。。。
	把最上面的盘子移动到目标柱。
*/
func hanota(A []int, B []int, C []int) []int {
	moveToTarget(len(A),&A,&B,&C)
	return C
}

//  把盘子移动到目标柱
func moveToTarget(size int,start ,center,target *[]int){
	// 递归的终止条件
	if size == 1{
		move(start,target)
		return
	}

	// 函数体和递推关系
	// 移动N-1个盘子到中间柱子
	moveToTarget(size-1,start,target,center)
	// 移动第N个盘子到C
	//move(start,target)
	moveToTarget(1,start,center,target)

	//  移动N-1个盘子到目标柱
	moveToTarget(size-1,center,start,target)
}
//  移动一个盘子到另一个柱子
func move(from,to *[]int){
	size:= len(*from)
	*to = append(*to, (*from)[size-1])
	*from = (*from)[:size-1]
}
/*
func hanota(A []int, B []int, C []int) []int {
	// 进行准备工作，对三个柱子的摆放顺序进行调整
	size := len(A)
	list := [3]*[]int{}
	list[0] = &A
	if size % 2 ==0 {
		list[1] = &B
		list[2] = &C
	}else{
		list[1] = &C
		list[2] = &B
	}
	//  记录最小的柱子的下标
	curPosition := 0
	//  循环进行接下来的两步操作
	for  len(C) != size {
		// 找到最小的柱子的下一个位置
		nextPosition := (curPosition+1)%3
		//  将这个柱子的盘子移动到下一个位置
		move(list[curPosition],list[nextPosition])

		prevPosition := curPosition
		curPosition  = nextPosition
		// 接着进行第二步操作：
		//当有一个柱子为空时，将盘子从非空的移动到空的柱子上，
		//当两个都非空时，把较小的那个移动较大哪一个柱子上。
		nextPosition = (curPosition + 1)%3

		prev := list[prevPosition]
		next := list[nextPosition]
		// 获取柱子最上层的盘子
		top1,top2:= math.MaxInt32,math.MaxInt32

		if length:=len(*prev);length!=0 {
			top1 = (*prev)[length-1]
		}
		if length:=len(*next);length!=0 {
			top2 = (*next)[length-1]
		}
		// 移动较小的到较大的上面
		if top1 > top2 {
			move(next,prev)
		}else if top1 < top2{
			move(prev,next)
		}
		//fmt.Printf("a=%#v,B=%#v,c=%#v\n",A,B,C)
	}
	return C
}

func move(from,to *[]int){
	size:= len(*from)
	*to = append(*to, (*from)[size-1])
	*from = (*from)[:size-1]
}
*/


//leetcode submit region end(Prohibit modification and deletion)


