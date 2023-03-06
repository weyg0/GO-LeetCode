package leetcode

/*func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	// 每个轮次的最大服务人数
	var gondolas []int
	t := len(customers)
	for i := 0; i < t-1; i++ {
		if customers[i] > 4 {
			gondolas = append(gondolas, 4)
			customers[i+1] += customers[i] - 4
		} else {
			gondolas = append(gondolas, customers[i])
		}
	}
	for customers[t-1] > 4 {
		gondolas = append(gondolas, 4)
		customers[t-1] -= 4
	}
	gondolas = append(gondolas, customers[t-1])
	// 截止到每个轮次的累计最大服务人数
	n := len(gondolas)
	for i := 1; i < n; i++ {
		gondolas[i] += gondolas[i-1]
	}
	// 计算最小轮转次数
	ans := -1
	max := 0
	for i := 0; i < n; i++ {
		tmp := gondolas[i]*boardingCost - (i+1)*runningCost
		if tmp > max {
			ans = i + 1
			max = tmp
		}
	}
	return ans
}*/

// 一次处理等待游客
func m1inOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	if runningCost >= boardingCost*4 {
		return -1
	}
	// 每个轮次的最大服务人数
	t := len(customers)
	for i := 0; i < t-1; i++ {
		if customers[i] > 4 {
			customers[i+1] += customers[i] - 4
			customers[i] = 4
		}
	}
	remain := 0
	if customers[t-1] > 4 {
		remain = customers[t-1] - 4
		customers[t-1] = 4
	}
	// 截止到每个轮次的累计最大服务人数
	for i := 1; i < t; i++ {
		customers[i] += customers[i-1]
	}
	// 计算最小轮转次数
	ans := -1
	max := 0
	for i := 0; i < t; i++ {
		tmp := customers[i]*boardingCost - (i+1)*runningCost
		if tmp > max {
			ans = i + 1
			max = tmp
		}
	}
	if remain == 0 {
		return ans
	}
	// 单独处理多余游客
	round := remain / 4
	last := remain % 4
	a := customers[t-1]*boardingCost - t*runningCost         // 原大小的利润
	b := round*4*boardingCost - round*runningCost            // 最后一轮不参与的利润
	c := (round*4+last)*boardingCost - (round+1)*runningCost // 最后一轮参与的利润
	if b >= c {
		if a+b > max {
			ans = t + round
		}
	} else {
		if a+c > max {
			ans = t + round + 1
		}
	}
	return ans
}

// 官方代码+一次处理
/*func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	// 纯赔
	if runningCost >= boardingCost*4 {
		return -1
	}
	// 处理原始轮转次数
	ans := -1
	wait, up, profit, max := 0, 0, 0, 0
	t := len(customers)
	for i := 0; i < t; i++ {
		wait += customers[i]
		if wait < 4 {
			up = wait
		} else {
			up = 4
		}
		wait -= up
		profit += up*boardingCost - runningCost
		if profit > max {
			max = profit
			ans = i + 1
		}
	}
	if wait == 0 {
		return ans
	}
	// 单独处理多余游客
	round, last := wait/4, wait%4
	if last*boardingCost <= runningCost {
		tmp := profit + round*4*boardingCost - round*runningCost
		if tmp > max {
			ans = t + round
		}

	} else {
		tmp := profit + (round*4+last)*boardingCost - (round+1)*runningCost
		if tmp > max {
			ans = t + round + 1
		}
	}
	return ans
}
*/
