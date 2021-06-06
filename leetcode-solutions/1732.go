func largestAltitude(gain []int) int {
    var answer int
    var cur int
    for _, v := range gain {
        k := cur + v
        if k > answer {
            answer = k
        }
        cur = k
    }
    return answer
}
