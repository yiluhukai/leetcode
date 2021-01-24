package  main

//在一排多米诺骨牌中，A[i] 和 B[i] 分别代表第 i 个多米诺骨牌的上半部分和下半部分。（一个多米诺是两个从 1 到 6 的数字同列平铺形成的 ——
//该平铺的每一半上都有一个数字。） 
//
// 我们可以旋转第 i 张多米诺，使得 A[i] 和 B[i] 的值交换。 
//
// 返回能使 A 中所有值或者 B 中所有值都相同的最小旋转次数。 
//
// 如果无法做到，返回 -1. 
//
// 
//
// 示例 1： 
//
// 
//
// 输入：A = [2,1,2,4,2,2], B = [5,2,6,2,3,2]
//输出：2
//解释：
//图一表示：在我们旋转之前， A 和 B 给出的多米诺牌。
//如果我们旋转第二个和第四个多米诺骨牌，我们可以使上面一行中的每个值都等于 2，如图二所示。
// 
//
// 示例 2： 
//
// 输入：A = [3,5,1,2,3], B = [3,6,3,3,4]
//输出：-1
//解释：
//在这种情况下，不可能旋转多米诺牌使一行的值相等。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= A[i], B[i] <= 6 
// 2 <= A.length == B.length <= 20000 
// 
// Related Topics 贪心算法 数组 
// 👍 62 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
	解法一:遍历整个多米诺骨牌，统计1-6每个数字在多米诺骨牌出现A、B、A|B中出现的次数
	只有在A｜B中出现的次数等于多米诺骨牌的数量的牌才是可以转成行相等的多米诺骨牌。
	然后根据A和B中出现的次数计算最小的旋转次数
	执行耗时:160 ms,击败了50.00% 的Go用户
	内存消耗:7.3 MB,击败了70.00% 的Go用户
	时间复杂度：统计1-6在多米诺骨牌中出现的次数  O(n)
	空间复杂度：O(1)
 */
//func minDominoRotations(A []int, B []int) int {
//	// 使用三个长度为6的数组统计多米诺骨牌中出现1-6在A、B、A｜B出现的次数
//	length:=len(A)
//	countA,countB,countAB := make([]int,6),make([]int,6),make([]int,6)
//
//	for i:=0;i<length;i++ {
//		countA[A[i]-1]++
//		countB[B[i]-1]++
//		countAB[A[i]-1]++
//		if A[i] != B[i] {
//			countAB[B[i]-1]++
//		}
//	}
//	// 遍历数组，看A｜B中是否有等于多米诺骨牌的数字
//
//	for i:=0;i<6;i++{
//		if countAB[i] == length{
//			// 表示可以旋转成行相等的多米诺骨牌
//			// 计算旋转次数
//			return minInt(length-countA[i],length-countB[i])
//		}
//	}
//	return -1
//}

/**
	贪心算法：通过局部的最优解得出最终的最优解(不一定适合所有问题)
	因为我们最终的目标是得到行相等的多米诺骨牌，而我们只能交换(旋转)A[i]和B[i].
	那么如果存在这样的行相等的多米诺骨牌，那么这个数字一定在A[0]和B[0]中，
	我们可以把A中全部变成A[0]或者B[0],或者把B中全部变成A[0]或者B[0].
	所以局部情况有四种。
	当我们可以将A中或者B中全部变成A[0],那么我们就没必要去检查B[0],
		1.若A[0] == B[0],检查就没有意义了，因为得到次数是一样的
		2.若A[0] != B[0],那么只有牌中只有A[0]或者B[0]的时候才可以，此时的旋转次数将和A[0]的最小旋转次数一样
	当不能A中或者B中不能全部变成A[0]，那么这个时候再去检查A或者B能否全部变成B[0]

	执行耗时:148 ms,击败了71.43% 的Go用户
	内存消耗:7.3 MB,击败了70.00% 的Go用户

	时间复杂度为O(n),
	空间复杂度为O(1),
	相比于上面的算法，这个算法比较的次数更少，当遇到A[0]和B[0]不相等的地方就停止了
 */
func minDominoRotations(A []int, B []int) int {
	rotateTimes := rotate(A,B,A[0])
	// 不能将A、B旋转为A[0]的行相等的多米诺骨牌
	if rotateTimes !=-1 {
		return rotateTimes
	}else{
		return rotate(A,B,B[0])
	}
}
/**
	将A或者B旋转后得到为target的行相等的多米诺骨牌的最少旋转次数，
	不能选装为-1
 */
func rotate(A,B []int,target int)int{
	// 统计target在A、B中出现的次数
	rotateCountA,rotateCountB:=0,0
	length:= len(A)
	for i:=0;i<length;i++{
		//如果出现了一个位置上A[i]和B[i]都不等于target.则说明无法旋转，返回-1。
		if A[i] != target && B[i]!=target{
			return -1
		}else if A[i] != target{
			rotateCountA++
		}else if B[i] != target{
			rotateCountB++
		}
	}
	// 在A或者B中都可以旋转成为A[0]的行相等多米诺骨牌，我们需要的是最小的旋转次数
	return minInt(rotateCountA,rotateCountB)
}
func minInt(a,b int)int{
	if a<b{
		return a
	}
	return b
}
//leetcode submit region end(Prohibit modification and deletion)


