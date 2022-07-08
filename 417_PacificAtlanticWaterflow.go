func dfs(heights [][]int, r int, c int, ocean *map[[2]int]bool, pastHeight int) {
	ROWS, COLS := len(heights), len(heights[0])

	if r < 0 || c < 0 || r >= ROWS || c >= COLS {
		return
	}

	if _, in := (*ocean)[[2]int{r, c}]; in {
		return
	}

	if heights[r][c] < pastHeight {
		return
	}

	(*ocean)[[2]int{r, c}] = true

	var directions [4][2]int = [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for _, d := range directions {
		dfs(heights, r+d[0], c+d[1], ocean, heights[r][c])
	}

	return

}

func pacificAtlantic(heights [][]int) [][]int {
	pacific := make(map[[2]int]bool)
	atlantic := make(map[[2]int]bool)

	ROWS, COLS := len(heights), len(heights[0])

	for c := 0; c < COLS; c++ {
		dfs(heights, 0, c, &pacific, 0)
		dfs(heights, ROWS-1, c, &atlantic, 0)
	}

	for r := 0; r < ROWS; r++ {
		dfs(heights, r, 0, &pacific, 0)
		dfs(heights, r, COLS-1, &atlantic, 0)
	}

	res := make([][]int, 0)

	for key, _ := range pacific {
		if _, in := atlantic[key]; in {
			res = append(res, []int{key[0], key[1]})
		}
	}

	return res

}
