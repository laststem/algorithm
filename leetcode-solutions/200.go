var A [][]bool
var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func numIslands(grid [][]byte) int {
    A = make([][]bool, len(grid))
    for i := 0; i < len(grid); i++ {
        A[i] = make([]bool, len(grid[i]))
    }
    var answer int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' && !A[i][j] {
                answer += f(grid, i, j)
            }
        }
    }
    return answer
}

func f(grid [][]byte, x, y int) int {
    A[x][y] = true
    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]
        if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || A[nx][ny] || grid[nx][ny] != '1' {
            continue
        }
        f(grid, nx, ny)
    }
    return 1
}
