package leetcode

func minSubarray(nums []int, p int) int {
	remainder := 0
	for _, n := range nums {
		remainder = (remainder + n) % p
	}
	if remainder == 0 {
		return 0
	}
	cur := 0                   // 前缀和
	last := map[int]int{0: -1} // 余数最后出现的位置
	ans := len(nums)
	for i, n := range nums {
		cur = (cur + n) % p
		key := (cur + p - remainder) % p
		if j, ok := last[key]; ok {
			if i-j < ans {
				ans = i - j
			}
		}
		last[cur] = i
	}
	if ans == len(nums) {
		return -1
	}
	return ans
}
