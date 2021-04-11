func largestNumber(nums []int) string {
    a := make([]string, len(nums))
    for i, v := range nums {
        a[i] = fmt.Sprintf("%d", v)
    }
    sort.Slice(a, func(i, j int) bool {
       return a[i]+a[j] > a[j]+a[i]
    })
    if a[0] == "0" {
        return "0"
    }
    var answer string
    for _, v := range a {
        answer += v
    }
    return answer
}
