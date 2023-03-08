package leetcode

import (
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	s := map[string]struct{}{} // GO语言中struct{}表示空struct，仅作为占位符，不占用任何内存
	dfs(expression, s)
	ans := make([]string, 0)
	for i := range s {
		ans = append(ans, i)
	}
	sort.Strings(ans)
	return ans
}

func dfs(exp string, s map[string]struct{}) {
	j := strings.Index(exp, "}")
	if j == -1 {
		s[exp] = struct{}{}
		return
	}
	i := strings.LastIndex(exp[:j], "{")
	a, c := exp[:i], exp[j+1:]
	for _, b := range strings.Split(exp[i+1:j], ",") {
		dfs(a+b+c, s)
	}
}
