package solutions

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		ans[i] = make([]int, n-2)
		for j := 0; j < n-2; j++ {
			for k := i; k <= i+2; k++ {
				for l := j; l <= j+2; l++ {
					if grid[k][l] > ans[i][j] {
						ans[i][j] = grid[k][l]
					}
				}
			}
		}
	}
	return ans
}
