func buildArray(target []int, n int) []string {
    k := 1
    var answer []string
    for i := 0; i < len(target); i++ {
        if (k == target[i]) {
            answer = append(answer, "Push")    
        } else {
            answer = append(answer, []string{"Push","Pop"}...)
            i--
        }
        k++
    }
    return answer
}
