/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start
import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
OuterLoop:
	for i := 1; i <= len(s)/2; i++ {
		if i == 1 {
			fmt.Println("Single str", i, len(s)/2)
			fmt.Println("Counts ", strings.Count(s, string(s[0])), " len", len(s))
			if strings.Count(s, string(s[0])) == len(s) {
				return true
			} else {
				continue
			}
		} else {
			fmt.Println("Current str length: ", i)
			if (len(s)-i)%i != 0 {
				continue
			} else {
				for j := 1; j < len(s)/i; j++ {
					fmt.Println("Two strings", s[j*i:], s[:i])
					if strings.HasPrefix(s[j*i:], s[:i]) {
						fmt.Println(s[j*i:])
						fmt.Println(s[:i])
						continue
					} else {
						continue OuterLoop
					}
				}
				return true
			}
		}
	}
	return false
}

// @lc code=end
