func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums := make([]int, 0)
    for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
        if i >= len(nums1) {
            nums = append(nums, nums2[j])
            j++
            continue
        }
        if j >= len(nums2) {
            nums = append(nums, nums1[i])
            i++
            continue
        }
        if nums1[i] < nums2[j] {
            nums = append(nums, nums1[i])
            i++
        } else {
            nums = append(nums, nums2[j])
            j++
        }
    }
    var answer float64
    if len(nums) % 2 > 0 {
        answer = float64(nums[(len(nums) / 2)])
    } else {
        k := float64(nums[len(nums)/2-1] + nums[len(nums)/2])
        answer = float64(k / 2.0)
    }
    return answer
}
