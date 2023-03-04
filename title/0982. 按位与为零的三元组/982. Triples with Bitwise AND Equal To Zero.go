package leetcode

// 暴力
/*func countTriplets(nums []int) int {
	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if nums[i]&nums[j]&nums[k] == 0 {
					ans++
				}
			}
		}
	}
	return ans
}*/

// 拆分枚举-哈希
/*func countTriplets(nums []int) int {
	count := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count[nums[i]&nums[j]]++
		}
	}
	ans := 0
	for k := 0; k < n; k++ {
		for ij, v := range count {
			if ij&nums[k] == 0 {
				ans += v
			}
		}
	}
	return ans
}*/

// 官方答案
/*func countTriplets(nums []int) int {
	arr := make([]int, 1<<16)
	MAX := 1<<16 - 1
	arr[0] = len(nums)
	for _, num := range nums {
		mask := MAX ^ num
		for i := mask; i > 0; i = (i - 1) & mask {
			arr[i]++
		}
	}
	ans := 0
	for _, a := range nums {
		for _, b := range nums {
			ans += arr[a&b]
		}
	}
	return ans
}*/

// 拆分枚举-数组
func countTriplets(nums []int) int {
	count := make([]int, 1<<16)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count[nums[i]&nums[j]]++
		}
	}
	ans := 0
	for k := 0; k < n; k++ {
		for ij, v := range count {
			if ij&nums[k] == 0 {
				ans += v
			}
		}
	}
	return ans
}
