package solutions

import "strconv"

/*func getFolderNames(names []string) []string {
	var ans []string
	nameMap := make(map[string]int)
	for _, name := range names {
		if val, ok := nameMap[name]; ok {
			newName := name + "(" + strconv.Itoa(val) + ")"
			_, existed := nameMap[newName]
			for existed {
				val++
				newName = name + "(" + strconv.Itoa(val) + ")"
				_, existed = nameMap[newName]
			}
			ans = append(ans, newName)
			nameMap[newName] = 1
			nameMap[name] = val + 1
		} else {
			ans = append(ans, name)
			nameMap[name] = 1
		}
	}
	return ans
}*/

// 原地修改
func getFolderNames(names []string) []string {
	nameMap := make(map[string]int)
	for i, name := range names {
		if val, ok := nameMap[name]; ok {
			newName := name + "(" + strconv.Itoa(val) + ")"
			_, existed := nameMap[newName]
			for existed {
				val++
				newName = name + "(" + strconv.Itoa(val) + ")"
				_, existed = nameMap[newName]
			}
			names[i] = newName
			nameMap[newName] = 1
			nameMap[name] = val + 1
		} else {
			nameMap[name] = 1
		}
	}
	return names
}
