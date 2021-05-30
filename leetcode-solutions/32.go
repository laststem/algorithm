func longestValidParentheses(s string) int {
    ans := 0
    stack := make([]int, 0)
    stack = append(stack, -1)
    for idx, c := range s {
        if c == '(' {
            stack = append(stack, idx)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                stack = append(stack, idx)
            } else {
                if ans < idx - stack[len(stack)-1] {
                    ans = idx - stack[len(stack)-1]
                }
            }
        }
    }
    return ans
}
