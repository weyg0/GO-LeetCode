package leetcode

import (
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	s := map[string]bool{}
	dfs(expression, s)
	ans := make([]string, 0)
	for i := range s {
		ans = append(ans, i)
	}
	sort.Strings(ans)
	return ans
}

func dfs(exp string, s map[string]bool) {
	j := strings.Index(exp, "}")
	if j == -1 {
		s[exp] = true
		return
	}
	i := strings.LastIndex(exp[:j], "{")
	a, c := exp[:i], exp[j+1:]
	for _, b := range strings.Split(exp[i+1:j], ",") {
		dfs(a+b+c, s)
	}
}
