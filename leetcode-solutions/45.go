func jump(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    ans := 1
    next := nums[0]
    jump := nums[0]
    for i := 1; i < len(nums); i++ {
        if i >= len(nums)-1 {
            break
        }
        if i+nums[i] > next {
            next = i+nums[i]
        }
        jump--
        if jump == 0 {
            ans++
            jump = next-i
        }
    }
    return ans
}
