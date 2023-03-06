package leetcode

import "strings"

// 由于最终状态为左侧都是'a'，右侧都是'b'，则可以遍历间隔，删除所有的左'b'和右'a'
/*func minimumDeletions(s string) int {
	leftB, rightA := 0, strings.Count(s, "a") // 间隔开始在最左侧
	ans := rightA
	for _, c := range s { // 遍历间隔
		if c == 'a' { // 发现左侧'a'，右'a'-1
			rightA--
		}
		if c == 'b' { // 发现左侧'b'，左'b'+1
			leftB++
		}
		if count := leftB + rightA; count < ans {
			ans = count
		}
	}
	return ans
}*/

// 优化掉if-else
func minimumDeletions(s string) int {
	del := strings.Count(s, "a") // 间隔开始在最左侧
	ans := del
	for _, c := range s { // 遍历间隔
		del += int(2*(c-'a') - 1) // 'a'则-1 'b'则+1
		if del < ans {
			ans = del
		}
	}
	return ans
}
