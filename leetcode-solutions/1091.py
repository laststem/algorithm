def shortestPathBinaryMatrix(grid: List[List[int]]) -> int:
    if grid[0][0] == 1 or grid[-1][-1]:
        return -1

    l = len(grid)-1
    d = [(i, j) for i in [-1, 0, 1] for j in [-1, 0, 1] if (i or j)]
    q = [(l+1, 1, l, l)]
    grid[l][l] = -1
    while q:
        _, cnt, x, y = heappop(q)
        if x == y == 0:
            return cnt
        cnt += 1
        for dx, dy in d:
            nx, ny = x + dx, y + dy
            include = 0 <= nx <= l and 0 <= ny <= l
            if include and (not grid[nx][ny] or grid[nx][ny] < -cnt):
                grid[nx][ny] = -cnt
                heappush(q, (cnt + max(nx, ny), cnt, nx, ny))
    return -1
