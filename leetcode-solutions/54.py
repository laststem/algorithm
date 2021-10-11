class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        answer = []
        n, m = len(matrix)-1, len(matrix[0])-1
        i, j = 0, 0
        while i <= n and j <= m:
            for y in range(j, m+1):
                answer.append(matrix[i][y])
            i += 1
            for x in range(i, n+1):
                answer.append(matrix[x][m])
            m -= 1
            if i <= n:
                y = m
                while y >= j:
                    answer.append(matrix[n][y])    
                    y -= 1
            n -= 1
            if j <= m:
                x = n
                while x >= i:
                    answer.append(matrix[x][j])
                    x -= 1
            j += 1
        return answer
