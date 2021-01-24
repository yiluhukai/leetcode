//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性： 
//
// 
// 每行中的整数从左到右按升序排列。 
// 每行的第一个整数大于前一行的最后一个整数。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 100 
// -104 <= matrix[i][j], target <= 104 
// 
// Related Topics 数组 二分查找 
// 👍 297 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/*
	暴力解法：遍历二维数组，和目标值对比，当等于目标值时，返回true
			解答成功:
					执行耗时:0 ms,击败了100.00% 的Go用户
					内存消耗:2.7 MB,击败了100.00% 的Go用户
	时间复杂度：O(n^2) 空间复杂度O(1)
*/
//func searchMatrix(matrix [][]int, target int) bool {
//	for _,arr :=range matrix{
//		for _,val := range arr {
//			if val == target{
//				return true
//			}
//		}
//	}
//	return false
//}

/*

	利用分置的思想查找元素: 1.现在整个矩阵的对角线上找，找到返回true
						2.当矩阵的对角线的下一个位置元素大于目标值或者到达矩阵的对角线末尾，
	开始在对角线的右上和左下区域继续查找。其他的两个区域肯定大于或者小于目标元素。
						3.当子区域中有一个区域找到目标元素，则返回true.负责返回false.
	4.当当前居于中的最小元素大于目标元素或者目标区域为空时，停止查找。
   执行耗时:4 ms,击败了86.17% 的Go用户
   内存消耗:2.8 MB,击败了9.46% 的Go用户

	时间复杂度:O(n) =

 */

//func searchMatrix(matrix [][]int, target int) bool {
//	//判断矩阵是否为空
//	colsLength:=len(matrix)
//
//	if colsLength == 0{
//		return false
//	}
//	rowLength := len(matrix[0])
//	if rowLength == 0{
//		return false
//	}
//	//在当前矩阵所在的区域查找
//	// 首次在整个矩阵中查找
//	return findMatrix(matrix,target,0,0,rowLength-1,colsLength-1)
//}

// 在矩阵的子区域中查找元素
// topX,topY,bottomX,bottomY 代表矩阵左上角和右下角的元素的坐标
// 注意元素坐标的计算

//func findMatrix(matrix [][]int, target ,topX,topY,bottomX,bottomY int)bool{
//
//	// 判断矩阵中的最小值是否大于目标元素 或者矩阵是否为空，当满足条件之一，就返回false
//	// 先判空，在判定是否小于矩阵的最小值
//	if (bottomX < topX  || bottomY < topY )|| matrix[topY][topX] > target{
//		return false
//	}
//	// 在矩阵的对角线上查找
//	diagonalLength:= bottomX-topX + 1
//
//	if length:= bottomY-topY+ 1;diagonalLength > length  {
//		diagonalLength = length
//	}
//	// 变量对角线上的值
//	for i:=0;i<diagonalLength;i++{
//		//对角线上找，找到返回true
//		if matrix[topY+i][topX+i] == target {
//			return true
//		}
//		//当矩阵的对角线的下一个位置元素大于目标值或者到达矩阵的对角线末尾，
//		//	开始在对角线的右上和左下区域继续查找。其他的两个区域肯定大于或者小于目标元素。
//		if i == diagonalLength-1 || matrix[topY+i+1][topX+i+1] > target {
//			return findMatrix(matrix,target,topX+i+1,topY,bottomX,topY+i) ||
//				findMatrix(matrix,target,topX,topY+i+1,topX+i,bottomY)
//		}
//	}
//	return false
//}


/*
	同样是利用分置的思想，上面的额每次查找都是去右上和左下两个区域查找的。我们可不可以只在一个区域中查找？
	1. 矩阵从上到下，从左到右是递增的。那么从左下到右上位置，矩阵一个方向是递增一个方向是递减的
	2. 我们使用两个指针分别指向矩阵的左下或者右上位置，若为左下，
	3. 当前指针指向的元素大于目标值，指针向上移动一个一位，
	4. 若当前位置的元素小于目标值，则指针想右移动。
	5. 然后在以当前指针指向的元素为一个顶点所构成的矩阵区域中继续查找，当有一个指针指向矩阵外，结束查找

    执行耗时:4 ms,击败了86.14% 的Go用户
    内存消耗:2.7 MB,击败了56.68% 的Go用户
	// 时间复杂度:O(m+n-1) ,最多比较m+n-1次
	// 空间复杂度  O(1),只有两个指针

	// 递归调用的深度log(m+n)  假设每次递归使用的空间复杂度为1
	那么总空间复杂度为2^0 + 2^1+....+2^(log(m+n)) = log(m+n)-1
	时间复杂度： log(m+n)*(m+n)
*/
func searchMatrix(matrix [][]int, target int) bool {
		//判断矩阵是否为空
		rowLength:=len(matrix)

		if rowLength == 0{
			return false
		}
		colsLength := len(matrix[0])
		if colsLength == 0{
			return false
		}
		// 从左下角的位置开始查找
		x,y:= 0,rowLength-1
		for x <= colsLength-1 && y >= 0{

			if val := matrix[y][x];val > target{
				// 指针上移
				y--
			}else if val == target{
				return true
			} else{
				// 小于目标值,指针右移
				x++
			}
		}
		return false
}


//leetcode submit region end(Prohibit modification and deletion)


