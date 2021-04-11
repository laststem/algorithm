type NumMatrix struct {
    nums [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    obj := NumMatrix{}
    obj.nums = make([][]int, len(matrix))
    for i := 0; i < len(matrix); i++ {
        obj.nums[i] = make([]int, len(matrix[i]))
        for j := 0; j < len(matrix[i]); j++ {
            obj.nums[i][j] = matrix[i][j]
            if i == 0 && j != 0 {
                obj.nums[i][j] += obj.nums[i][j-1]
            } else if j == 0 && i != 0 {
                obj.nums[i][j] += obj.nums[i-1][j]
            } else if i != 0 && j != 0 {
                obj.nums[i][j] = matrix[i][j] - obj.nums[i-1][j-1] + obj.nums[i-1][j] + obj.nums[i][j-1]
            }
        }
    }
    return obj
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    answer := this.nums[row2][col2]
    if row1-1 >= 0 {
        answer -= this.nums[row1-1][col2]
    }
    if col1-1 >= 0 {
        answer -= this.nums[row2][col1-1]
    }
    if row1-1>=0&&col1-1>=0 {
        answer += this.nums[row1-1][col1-1]
    }
    return answer
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
