func canConstruct(ransomNote string, magazine string) bool {
    characters := make(map[rune]int, 0)
    for _, c := range magazine {
        characters[c]++
    }
    for _, c := range ransomNote {
        if characters[c] > 0 {
            characters[c]--
        } else {
            return false
        }
    }
    return true
}
