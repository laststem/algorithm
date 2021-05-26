func removeElement(nums []int, val int) int {
    j := 0
    for i := 0; i < len(nums); i++ {
        for i < len(nums) && nums[i] == val {
            i++
        }
        if i == len(nums) {
            break
        }
        nums[j] = nums[i]
        j++
    }
    return j
}
