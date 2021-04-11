func isValidSerialization(preorder string) bool {
    splat := strings.Split(preorder, ",")
    
    c := 1
    for _, v := range splat {
        if c-1 < 0 {
            return false
        }
        c -= 1
        if v != "#" {
            c += 2
        }
    }
    return c==0
}
