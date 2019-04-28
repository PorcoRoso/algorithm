package greedy

import "sort"

// 455.Assign Cookies
// g represents children content factor
// s represents cookies   
func findContentChildren(g []int, s []int) int {
	sum := 0
	i, j := 0, 0

	sort.Ints(g)
	sort.Ints(s)
	for ; i < len(s) && j < len(g); {
		for ; j < len(g) && i < len(s); {
			if s[i] < g[j] {
				i++
			} else {
				sum++
				i++
				j++
			}
		}
	}

	return sum
}
