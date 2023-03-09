package leetcode

import "strings"

// 暴力
/*func minimumRecolors(blocks string, k int) int {
	ans := len(blocks)
	for i := 0; i < len(blocks)-k+1; i++ {
		tmp := 0
		for j := i; j < k+i; j++ {
			if blocks[j] == 'W' {
				tmp++
			}
		}
		if tmp < ans {
			ans = tmp
		}
	}
	return ans
}*/

// 滑动窗口
func minimumRecolors(blocks string, k int) int {
	cnt := strings.Count(blocks[:k], "W")
	ans := cnt
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			cnt++
		}
		if blocks[i-k] == 'W' {
			cnt--
		}
		if cnt < ans {
			ans = cnt
		}
	}
	return ans
}
