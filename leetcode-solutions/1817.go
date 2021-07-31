func findingUsersActiveMinutes(logs [][]int, k int) []int {
    m := make(map[int]map[int]struct{})
    for _, log := range logs {
        _, ok := m[log[0]]
        if !ok {
            m[log[0]] = make(map[int]struct{})
        }
        m[log[0]][log[1]] = struct{}{}
    }
    answer := make([]int, k)
    for _, maps := range m {
        answer[len(maps)-1]++
    }
    return answer
}
