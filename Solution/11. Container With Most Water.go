package Solution

func maxArea(height []int) int {
	l, r, max := 0, len(height)-1, 0
	for l < r {
		w := r - l
		h := 0
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}
		tmp := h * w
		if tmp > max {
			max = tmp
		}
	}
	return max
}
