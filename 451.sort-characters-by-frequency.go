/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
import (
	"fmt"
	"sort"
)

type kv struct {
	Key   string
	Value int
}

func frequencySort(s string) string {
	m := make(map[string]int)
	var ss []kv
	for i := 0; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok {
			m[string(s[i])]++
		} else {
			m[string(s[i])] = 1
		}
	}
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	res := ""
	for _, kv := range ss {
		for index := 0; index < kv.Value; index++ {
			res += kv.Key
		}
	}
	fmt.Println(ss)
	return res
}

// @lc code=end
