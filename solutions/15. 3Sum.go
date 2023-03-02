package solutions

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 { // 最左侧指针指向的数字为0，则三数之和不可能为0
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[l] + nums[i] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				ans = append(ans, []int{nums[l], nums[i], nums[r]})
				for l < r && nums[l] == nums[l+1] { // 去重
					l++
				}
				for l < r && nums[r] == nums[r-1] { // 去重
					r--
				}
				l++
				r--
			}
		}
	}
	return ans
}
