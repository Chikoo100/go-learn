func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}


func maxArea(height []int) int {
    res := 0
    l, r := 0, len(height) - 1
    
    for l < r {
        
        if height[l] < height[r] {
            area := height[l] * (r - l)
            res = max(res, area)
            l += 1
        } else {
            area := height[r] * (r - l)
            res = max(res, area)
            r -= 1
        }
    }
    return res
    
}
