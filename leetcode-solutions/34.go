func searchRange(nums []int, target int) []int {
    first, last := f(nums, target), f2(nums, target)
    return []int{first, last}
}

func f(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    mid := (lo+hi)/2
    ans := -1
    for lo <= hi {
        mid = (lo+hi)/2
        if nums[mid] >= target {
            hi = mid-1
            if nums[mid] == target {
                ans = mid
            }
        } else {
            lo = mid+1
        }
    }
    return ans
}

func f2(nums []int, target int) int {
    lo, hi := 0, len(nums)-1
    mid := (lo+hi)/2
    ans := -1
    for lo <= hi {
        mid = (lo+hi)/2
        if nums[mid] <= target {
            lo = mid+1
            if nums[mid] == target {
                ans = mid
            }
        } else {
            hi = mid-1
        }
    }
    return ans
}
