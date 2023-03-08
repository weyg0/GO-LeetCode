## 剑指 Offer 47. 礼物的最大价值

#### 解法1 DP

定义$dp[i][j]$为从左上角起点走到$(i-1,j-1)$的最大价值，动态规划转移方程如下：
$$
dp[i][j]=\max(dp[i-1][j],dp[i][j-1])+grid[i-1][j-1]
$$

```go
func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 解法2 DP+空间优化

直接使用原$gird$数组，定义$grid[i][j]$为从左上角起点走到$(i,j)$的最大价值，动态规划转移方程如下：
$$
grid[i][j]=\max(grid[i-1][j],grid[i][j-1])+grid[i][j]
$$

```go
func maxValue(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
```

