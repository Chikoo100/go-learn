func twoSum(nums []int, target int) []int {
    var visited = make(map[int]int)
    
    for key, val := range nums {
        if _, member := visited[val]; member {
            return []int{visited[val], key}
        }
        visited[target - val] = key
    }
    
    return []int{}
}
