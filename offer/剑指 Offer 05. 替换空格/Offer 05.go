package offer

// 调库
/*func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}*/

// 调库
/*func replaceSpace(s string) string {
	return strings.Join(strings.Split(s, " "), "%20")
}*/

func replaceSpace(s string) string {
	ans := ""
	for _, c := range s {
		if c == ' ' {
			ans += "%20"
		} else {
			ans += string(c)
		}
	}
	return ans
}
