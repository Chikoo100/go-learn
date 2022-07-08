func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    charCount := make(map[byte]int)
    
    for i := 0; i < len(s); i++ {
        charCount[s[i]] += 1
        charCount[t[i]] -= 1
    }
    
    for _, v := range charCount {
        if v != 0 {
            return false
        }
    }
    
    return true
    
}
