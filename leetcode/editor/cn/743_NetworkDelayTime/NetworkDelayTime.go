package  main

import "math"

//有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， w
//i 是一个信号从源节点传递到目标节点的时间。 
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：times = [[1,2,1]], n = 2, k = 1
//输出：1
// 
//
// 示例 3： 
//
// 
//输入：times = [[1,2,1]], n = 2, k = 2
//输出：-1
// 
//
// 
//
// 提示： 
//
// 
// 1 <= k <= n <= 100 
// 1 <= times.length <= 6000 
// times[i].length == 3 
// 1 <= ui, vi <= n 
// ui != vi 
// 0 <= wi <= 100 
// 所有 (ui, vi) 对都 互不相同（即，不含重复边） 
// 
// Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 
// 👍 408 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
// 该题是求k到所有图的所有顶点中最小权值和
// 使用深度优先搜索去统计k到所有图的所有顶点中最小;
//权值和
//解答成功:
//执行耗时:120 ms,击败了5.34% 的Go用户
//内存消耗:7.5 MB,击败了20.58% 的Go用户
// 时间复杂度： 创建邻接表 O(E) + O(N * E) 遍历每个节点的所有邻接点
// 空间复杂度	邻接表 O(E) + dist O(N)
/*
func networkDelayTime(times [][]int, n int, k int) int {
	// 还是将上面的输入times转为图的邻接表的形式
	// time = [2,1,1],
	adjList := make([][][]int, n + 1)
	for _, time :=range times{
		start,end,weight := time[0],time[1],time[2]
		adjList[start] = append(adjList[start],[]int{end,weight})
	}
	//有 n 个网络节点，标记为 1 到 n。
	// 存储k到每个顶点m的最小权值和
	dist := make([]int,n+1)
	// 初始化为 无穷大
	for index :=range dist{
		dist[index] = math.MaxInt64
	}
	// 深度遍历图
	dist[k] = 0
	dfsGraph(adjList, k ,dist)

	// 查找dist最大值
	maxWeight := 0
	length :=len(dist)
	for  i:=1;i <length; i++{
		v := dist[i]
		if v == math.MaxInt64 {
			return -1
		}
		maxWeight =  maxInt( maxWeight , v)
	}
	return maxWeight
}
func maxInt(a,b int )int{
	if a > b{
		return a
	}
	return b
}


func dfsGraph(adjList [][][]int,k int,dist []int){
	// 遍历顶点k的邻接点
	for _, edge := range adjList[k] {

		// edge[0] 邻接点   edge[1] 是 k到点vertex[0]的weight
		vertex,weight := edge[0],edge[1]
		if totalWeight:= dist[k] + weight ;totalWeight < dist[vertex]{
			// 更新 k到该顶点最小权值和
			dist[vertex] = totalWeight
			dfsGraph(adjList,vertex,dist)
		}
		// 当weight >= dist[vertex[0]], 没有必要去修改后续的权重值，这个过程称之为剪枝
	}
}*/
// 狄杰斯科拉算法 去计算单源最短路径
// 狄杰斯科拉算法：类似贪心算法和广度优先遍历。每次都去求第i近的路径

//解答成功:
//执行耗时:60 ms,击败了45.37% 的Go用户
//内存消耗:7.6 MB,击败了14.96% 的Go用户
// 时间复杂度： 创建邻接表 O(E) + O(N^2)查找当前顶点的所有临接点中距离k最近的下一个点
// 空间复杂度	邻接表 O(E) + dist 和 visited O(V)
func networkDelayTime(times [][]int, n int, k int) int {
	// 还是将上面的输入times转为图的邻接表的形式
	// time = [2,1,1],
	adjList := make([][][]int, n + 1)
	for _, time :=range times{
		start,end,weight := time[0],time[1],time[2]
		adjList[start] = append(adjList[start],[]int{end,weight})
	}
	// 存储k到每个顶点m的最小权值和
	dist := make([]int,n+1)
	// 初始化为 无穷大
	for index :=range dist{
		dist[index] = math.MaxInt64
	}
	// 深度遍历图
	visited := make([]bool, n+1 )
	dist[k] = 0
	// 存放离k第i近的顶点
	vertexes := make([]int, 0,n+1)
	for {
		vertex:= -1
		minWeight := math.MaxInt64
		for i:=1;i < n+1;i++{
			// 查找当前顶点的所有临接点中距离k最近的下一个点
			if !visited[i] && dist[i] < minWeight{
				minWeight = dist[i]
				vertex = i
			}
		}
		// 找不到离k最近且没有访问的顶点
		if vertex == -1 {
			break
		}
		// 将该顶点放入到 vertexes
		vertexes = append(vertexes, vertex)
		visited[vertex] = true
		for _,edge := range adjList[vertex]{
			v,weight := edge[0],edge[1]
			// 更新k到vertex的临界点的距离
			dist[v] = int(math.Min( float64(dist[v]),float64(dist[vertex] + weight)))
		}
	}
	var maxWeight float64
	for  i:=1;i < n+1; i++{
		weight := dist[i]
		if weight == math.MaxInt64 {
			return -1
		}
		maxWeight =  math.Max( maxWeight , float64(weight))
	}
	return int(maxWeight)
}

//leetcode submit region end(Prohibit modification and deletion)

func main(){
	times := [][]int { {2,1,1},{2,3,1},{3,4,1}}
	networkDelayTime(times,4,2)
}

