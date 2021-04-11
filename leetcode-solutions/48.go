func rotate(matrix [][]int)  {
    maxLen := len(matrix)
    for i := 0; i < maxLen/2; i++ {
        for j := i; j < maxLen-i-1; j++ {
            a := matrix[i][j]
            matrix[i][j] = matrix[maxLen-j-1][i]
            matrix[maxLen-j-1][i] = matrix[maxLen-i-1][maxLen-j-1]
            matrix[maxLen-i-1][maxLen-j-1] = matrix[j][maxLen-i-1]
            matrix[j][maxLen-i-1] = a
        }
    }
}
