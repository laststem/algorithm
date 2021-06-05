func findMin(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    } 
    lo, hi := 0, len(nums)-1
    if nums[lo] < nums[hi] {
        return nums[lo]
    }
    
    target := nums[hi]
    for lo <= hi {
        mid := (lo+hi)/2
        if nums[mid] < target {
            target = nums[mid]
            hi = mid-1
        } else {
            lo = mid+1
        }
    }
    return target
}
