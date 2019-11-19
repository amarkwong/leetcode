/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
import (
	"math"
)

func pIncrement(p int) int {
	if p == 2 {
		return 3
	} else {
		return p + 2
	}
}

func countPrimes(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 0
	case 2:
		return 0
	case 3:
		return 1
	case 4:
		return 2
	case 5:
		return 2
	default:
		var slice = make([]bool, n)
		var ns = math.Sqrt(float64(n))
		var nsqrt = int(ns)
		for p := 2; p <= nsqrt; p = pIncrement(p) {
			for i := p * p; i < n; i += p {
				slice[i] = true
			}
		}
		var cnt int = 0
		for j := 0; j < n; j++ {
			if !slice[j] {
				cnt++
			}
		}
		return cnt - 2
	}
}

// @lc code=end
