func intersect(nums1 []int, nums2 []int) []int {
    m := make(map[int]int, 0)
    for _, k := range nums1 {
        m[k]++;
    }
    var answer []int
    for _, k := range nums2 {
        if m[k] > 0 {
            answer = append(answer, k)
            m[k]--
        }
    }
    return answer
}
