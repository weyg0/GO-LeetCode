package interview

func findLongestSubarray(array []string) []string {
	visited := map[int]int{0: -1}
	sum, max, start := 0, 0, 0
	for right, s := range array {
		if s[0] >= 'A' {
			sum++
		} else {
			sum--
		}
		if left, ok := visited[sum]; !ok {
			visited[sum] = right
		} else if right-left > max {
			max = right - left
			start = left + 1
		}
	}
	return array[start : start+max]
}
