func groupAnagrams(strs []string) [][]string {
    index := make(map[[26]int][]string)
    
    for _, word := range strs {
        var key [26]int
        for _, char := range word {
            charVal := int(char) - int('a')
            key[charVal] += 1
        }
        index[key] = append(index[key], word)
    }
    
    res := make([][]string, 0)
    
    for _, val := range index {
        res = append(res, val)
    }
    
    return res
    
}
