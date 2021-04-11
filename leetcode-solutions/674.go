func findLengthOfLCIS(nums []int) int {
    l := len(nums)
    if l <= 1 {
        return l
    }
    
    var answer int
    pre := nums[0]
    x := 1
    for i:=1; i<len(nums); i++ {
        if nums[i] > pre {
          x++  
        } else {
            answer = max(answer, x)
            x = 1
        }
        pre = nums[i]
    }
    return max(answer, x)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
