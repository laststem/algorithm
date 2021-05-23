func removeDuplicates(nums []int) int {
    length := len(nums)
    if length <= 1 {
        return length
    }
    j := 0
    for i := 0; i < length; i++ {
        for i+1 < length && nums[i] == nums[i+1] {
            i++
        }
        nums[j] = nums[i]
        j++
    }
    return j
}
