var D []int

func rob(nums []int) int {
    D = make([]int, len(nums))
    for i := 0; i<len(nums); i++ {
        D[i] = -1
    }
    return f(nums, 0)
}

func f(nums []int, cur int) int {
    if cur >= len(nums) {
        return 0
    }
    if D[cur] != -1 {
        return D[cur]
    }
    D[cur] = f(nums, cur+1)
    x := f(nums, cur+2) + nums[cur]
    if D[cur] < x {
        D[cur] = x
    }
    return D[cur]
}
