func containsDuplicate(nums []int) bool {
    
    visited := make(map[int]bool, len(nums))
    
    for _, val := range nums {
        if _, member := visited[val]; member {
            return true
        }
        visited[val] = true
    }
    return false 
}
