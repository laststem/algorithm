func generateParenthesis(n int) []string {
    answer := []string{}
    f(n, n, "", &answer)
    return answer
}

func f(x, y int, s string, answer *[]string) {
    if x == 0 && y == 0 {
        *answer = append(*answer, s)
        return
    }
    if x > 0 {
        f(x-1, y, s+"(", answer)
    }
    if x < y {
        f(x, y-1, s+")", answer)
    }
}
