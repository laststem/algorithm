func allPathsSourceTarget(graph [][]int) [][]int {
    answer := new([][]int)
    return dfs(answer, graph, 0, len(graph) - 1, []int{0}, map[int]bool{0:true})
}

func dfs(answer *[][]int, graph [][]int, current, target int, path []int, visit map[int]bool) [][]int {
    if current == target {
        p := make([]int, len(path))
        copy(p, path)
        *answer = append(*answer, p)
        return *answer
    }
    
    for _, node := range graph[current] {
        if !visit[node] {
            visit[node] = true
            path = append(path, node)
            dfs(answer, graph, node, target, path, visit)
            path = path[:len(path)-1]
            visit[node] = false
        }
    }
    return *answer
}
