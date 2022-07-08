func rob(nums []int) int {
    
    dp := make([]int, len(nums))
    
    max := func(a int, b int) int {
        if a > b {
            return a
        }
        return b
    }
    
    for i, val := range nums {
        if i == 0 {
            dp[i] = val
            continue
        }
        
        if i == 1 {
            dp[i] = max(dp[0], val)
            continue
        }
        dp[i] = max(dp[i-1], dp[i-2] + val)
        
    }
    
    return dp[len(nums) - 1]
    
}
