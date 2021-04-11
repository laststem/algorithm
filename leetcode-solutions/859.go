func buddyStrings(A string, B string) bool {
    if len(A) != len(B) {
        return false
    }
    
    if A == B {
        m := make(map[rune]bool, 0)
        for _, c := range A {
            if m[c] {
                return true
            }
            m[c] = true
        }
        return false
    }
    a, b := -1, -1
    for i := 0; i < len(A); i++ {
        if A[i] != B[i] {
            if a == -1 {
                a = i
            } else if b == -1 {
                b = i
            } else {
                return false
            }
        }
    }
    return a != -1 && A[a] == B[b] && A[b] == B[a]
}
