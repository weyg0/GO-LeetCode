package leetcode

/*func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	gapEnergy, gapExperience := 0, 0
	for i := 0; i < len(experience); i++ {
		if initialExperience <= experience[i] {
			gapExperience += experience[i] + 1 - initialExperience
			initialExperience = experience[i] + 1
		}
		initialExperience += experience[i]
	}
	for i := 0; i < len(energy); i++ {
		initialEnergy -= energy[i]
	}
	if initialEnergy > 0 {
		gapEnergy = 0
	} else {
		gapEnergy = 1 - initialEnergy
	}
	return gapEnergy + gapExperience
}*/

// 简洁
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	ans := 0
	for i := 0; i < len(energy); i++ {
		if initialEnergy <= energy[i] {
			ans += energy[i] + 1 - initialEnergy
			initialEnergy = energy[i] + 1
		}
		initialEnergy -= energy[i]
		if initialExperience <= experience[i] {
			ans += experience[i] + 1 - initialExperience
			initialExperience = experience[i] + 1
		}
		initialExperience += experience[i]
	}
	return ans
}
