func numsSameConsecDiff(n int, k int) []int {
    var q []int
    for i := 1; i <= 9; i++ {
        q = append(q, i)
    }
    n--
    for n > 0 {
        l := len(q)
        for _, no := range q {
            c := no % 10
            for i := 0; i <= 9; i++ {
                if abs(c-i) == k {
                    q = append(q, no*10+i)
                } 
            }
        }
        q = q[l:]
        n--
    }
    return q
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}
