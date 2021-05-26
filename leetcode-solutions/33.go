func search(nums []int, target int) int {
    origin := idx(nums)
    if nums[origin] <= target && target <= nums[len(nums)-1] {
        return f(nums, origin, len(nums)-1, target)
    }
    return f(nums, 0, origin-1, target)
}

func idx(nums []int) int {
    lo, hi := 0, len(nums)-1
    for lo < hi {
        m := (lo + hi) / 2
        if nums[m] < nums[hi] {
            hi = m
        } else {
            lo = m+1
        }
    }
    return lo
}
                  
func f(nums []int, lo, hi, val int) int {
    for lo <= hi {
        m := (lo + hi) / 2
        if nums[m] < val {
            lo = m+1
        } else if nums[m] == val {
            return m
        } else {
            hi = m-1
        }
    }
    return -1
}
                
