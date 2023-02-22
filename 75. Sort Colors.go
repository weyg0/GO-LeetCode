package leetCode

// 内置函数
/*import (
	"fmt"
	"sort"
)

func sortColors(nums []int) {
	sort.Ints(nums)
	fmt.Println(nums)
}*/

// 计数排序
/*func sortColors(nums []int) {
	n := len(nums)
	bucket := make([]int, 3)
	for i := 0; i < n; i++ {
		bucket[nums[i]] += 1
	}
	index := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < bucket[i]; j++ {
			nums[index] = i
			index++
		}
	}
}*/

// 刷油漆
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		nums[i] = 2
		if n < 2 {
			nums[one] = 1
			one++
		}
		if n < 1 {
			nums[zero] = 0
			zero++
		}
	}
}
