package leetcode

func minSubarray(nums []int, p int) int {
	total := 0
	for _, v := range nums {
		total = (total + v) % p
	}
	if total == 0 {
		return 0
	}
	last := map[int]int{0: -1}
	n := len(nums)
	ans := n
	cur := 0
	for i, v := range nums {
		cur = (cur + v) % p
		target := (cur - total + p) % p
		if key, ok := last[target]; ok {
			if i-key < ans {
				ans = i - key
			}
		}
		last[cur] = i
	}
	if ans == n {
		return -1
	}
	return ans
}
