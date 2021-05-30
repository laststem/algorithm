func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    needleLen := len(needle)
    for i := 0; i+needleLen <= len(haystack); i++ {
        cur := haystack[i:i+needleLen]
        if cur == needle {
            return i
        }
    }
    return -1
}
