package main

//节点间通路。给定有向图，设计一个算法，找出两个节点之间是否存在一条路径。
//
// 示例1:
//
//  输入：n = 3, graph = [[0, 1], [0, 2], [1, 2], [1, 2]], start = 0, target = 2
// 输出：true
//
//
// 示例2:
//
//  输入：n = 5, graph = [[0, 1], [0, 2], [0, 4], [0, 4], [0, 1], [1, 3], [1, 4], [
//1, 3], [2, 3], [3, 4]], start = 0, target = 4
// 输出 true
//
//
// 提示：
//
//
// 节点数量n在[0, 1e5]范围内。
// 节点编号大于等于 0 小于 n。
// 图中可能存在自环和平行边。
//
// Related Topics 深度优先搜索 广度优先搜索 图
// 👍 38 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 图的广度优先搜索 节点是否连通
//解答成功:
//执行耗时:148 ms,击败了92.31% 的Go用户
//内存消耗:22 MB,击败了60.44% 的Go用户
// G(V,E)
// 时间复杂度：建立邻接表的时间复杂度为 O(|E|) + 查找节点的过程O(|V|)
// 空间复杂度：建立邻接表 O(|V| + |E|) + O(V) 是否访问顶点的标记数组 + O(V) 存储每层定点的数组
/*
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	// 将入上面的输入转换成一个邻接表
	// 使用一个二维切片作为邻接表
	adjList := make([][]int , n)
	// edge 是有顶部表示的两条边，
	for _,edge := range graph {
		startVertex,endVertex := edge[0],edge[1]
		adjList[startVertex] = append(adjList[startVertex], endVertex)
	}
	// 以start为起点图的广度优先搜索
	// 使用visited去存储那个定点是否被访问过了，避免死循环
	// vertexes存储存储每层的临界点
	visited,vertexes := make([]bool,n),make([]int,0,n)
	index :=  0
	vertexes = append(vertexes,start)
	visited[start] = true
	for  len(vertexes) > index  {
		// 从vertexes中取出临界点
		vertex := vertexes[index]
		index++
		if vertex == target {
			return true
		}
		// 找出这个节点的所有非访问邻接点

		for _,v := range adjList[vertex] {
			if !visited[v] {
				visited[v] = true
				vertexes = append(vertexes, v)
			}
		}
	}
	return false
}
*/
// 使用深度优先遍历
//解答成功:
//执行耗时:144 ms,击败了92.31% 的Go用户
//内存消耗:23.8 MB,击败了29.67% 的Go用户
// 时间复杂度 ： 建立邻接表的时间复杂度为 O(|E|) + 查找节点的过程O(|V|)
// 空间复杂度：建立邻接表 O(|V| + |E|) + O(V) 是否访问顶点的标记数组
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	// 将入上面的输入转换成一个邻接表
	// 使用一个二维切片作为邻接表
	adjList := make([][]int, n)
	// edge 是有顶部表示的两条边，
	for _, edge := range graph {
		startVertex, endVertex := edge[0], edge[1]
		adjList[startVertex] = append(adjList[startVertex], endVertex)
	}
	// 以start为起点 图的深度优先搜索
	// 使用visited去存储那个定点是否被访问过了，避免死循环
	// vertexes存储存储每层的临界点
	visited := make([]bool, n)

	// 找出这个节点的所有非访问邻接点

	return dfsGraph(adjList, start, target, visited)
}

func dfsGraph(adjList [][]int, start, target int, visited []bool) bool {
	if start == target {
		return true
	}
	visited[start] = true
	for _, v := range adjList[start] {
		if !visited[v] {
			if dfsGraph(adjList, v, target, visited) {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

// 使用深度优先遍历

func main() {

	graph := [][]int{{0, 1}, {0, 3}, {0, 6}, {0, 7}, {0, 28}, {1, 2}, {2, 44}, {2, 59}, {3, 4}, {3, 13},
		{3, 14}, {3, 17}, {4, 5}, {4, 8}, {4, 10}, {4, 11}, {5, 16}, {5, 20}, {5, 25}, {6, 29}, {6, 36}, {6, 41}, {7, 9}, {7, 21}, {7, 37}, {9, 12}, {9, 19}, {9, 30}, {10, 39}, {12, 18}, {12, 22}, {12, 34}, {13, 32}, {14, 15}, {14, 35}, {15, 24}, {15, 51}, {16, 23}, {16, 27}, {17, 61}, {20, 31}, {21, 63}, {22, 33}, {23, 42}, {24, 26}, {24, 40}, {25, 60}, {27, 46}, {29, 48}, {29, 49}, {30, 45}, {30, 53}, {31, 38}, {33, 47},
		{33, 55}, {33, 57}, {34, 52}, {37, 58}, {39, 43}, {39, 44}, {43, 62}, {46, 56}, {47, 50}, {51, 54}}
	findWhetherExistsPath(64, graph, 0, 63)
}
