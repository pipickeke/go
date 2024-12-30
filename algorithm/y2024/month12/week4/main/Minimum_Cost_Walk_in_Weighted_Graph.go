package main

/**
题目：3108. 带权图里旅途的最小代价

给你一个 n 个节点的带权无向图，节点编号为 0 到 n - 1 。

给你一个整数 n 和一个数组 edges ，其中 edges[i] = [ui, vi, wi] 表示节点 ui 和 vi 之间有一条权值为 wi 的无向边。

在图中，一趟旅途包含一系列节点和边。旅途开始和结束点都是图中的节点，且图中存在连接旅途中相邻节点的边。注意，一趟旅途可能访问同一条边或者同一个节点多次。

如果旅途开始于节点 u ，结束于节点 v ，我们定义这一趟旅途的 代价 是经过的边权按位与 AND 的结果。换句话说，如果经过的边对应的边权为 w0, w1, w2, ..., wk ，那么代价为w0 & w1 & w2 & ... & wk ，其中 & 表示按位与 AND 操作。

给你一个二维数组 query ，其中 query[i] = [si, ti] 。对于每一个查询，你需要找出从节点开始 si ，在节点 ti 处结束的旅途的最小代价。如果不存在这样的旅途，答案为 -1 。

返回数组 answer ，其中 answer[i] 表示对于查询 i 的 最小 旅途代价。

*/

func minimumCost(n int, edges, query [][]int) []int {
	type edge struct {
		to, w int
	}

	g := make([][]edge, n)
	for _, e := range edges {
		x, y, w := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, w})
		g[y] = append(g[y], edge{x, w})
	}

	ring := make([]int, n)
	for i := range ring {
		ring[i] = -1
	}

	ringAndSum := []int{}

	var dfs func(int) int
	dfs = func(x int) int {
		ring[x] = len(ringAndSum)
		and := -1
		for _, e := range g[x] {
			and &= e.w
			if ring[e.to] < 0 {
				and &= dfs(e.to)
			}
		}
		return and
	}

	for i, id := range ring {
		if id < 0 {
			ringAndSum = append(ringAndSum, dfs(i))
		}
	}

	ans := make([]int, len(query))
	for i, q := range query {
		s, t := q[0], q[1]
		if ring[s] != ring[t] {
			ans[i] = -1
		} else {
			ans[i] = ringAndSum[ring[s]]
		}
	}
	return ans
}
