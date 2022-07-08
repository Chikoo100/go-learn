func dfs(grid [][]byte, r int, c int, visited map[[2]int]bool) bool {
    ROWS, COLS := len(grid), len(grid[0])
    if r < 0 || r >= ROWS || c < 0 || c >= COLS {
        return false
    }
    
    if _, ok := visited[[2]int{r, c}]; ok {
        return false
    }
    
    if grid[r][c] == '0' {
        return false
    }
    
    visited[[2]int{r, c}] = true
    
    var directions [4][2]int = [4][2]int{{1,0}, {0,1}, {-1,0}, {0,-1}}
    
    for _, d := range directions {
        dfs(grid, r + d[0], c + d[1], visited)
    }
    
    return true
    
}

func numIslands(grid [][]byte) int {
    
    ROWS, COLS := len(grid), len(grid[0])
    visited := make(map[[2]int]bool)
    res := 0
    
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if grid[r][c] == '1' && dfs(grid, r, c, visited){
                res += 1
            }   
        }
    }
    return res 
}
