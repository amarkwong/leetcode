import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */

// @lc code=start

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	// FIXME: What do we do if the stack is empty, though?
	l := len(s)
	if l == 0 {
		return nil, 0
	}
	return s[:l-1], s[l-1]
}

func calPoints(ops []string) int {
	st := make(stack, 0)
	sum, v1, v2 := 0, 0, 0
	for _, char := range ops {
		num, _ := strconv.Atoi(char)
		if num >= -30000 && num <= 30000 && char != "C" && char != "D" && char != "+" {
			sum += num
			st = st.Push(num)
			fmt.Println(st)
		}
		if char == "C" {
			st, v1 = st.Pop()
			sum -= v1
			fmt.Println("Pop: ", v1, " Sum: ", sum)
		}
		if char == "+" {
			st, v1 = st.Pop()
			st, v2 = st.Pop()
			sum += v1 + v2
			st = st.Push(v2)
			st = st.Push(v1)
			st = st.Push(v1 + v2)
			fmt.Println("Pop: ", v1, " Sum: ", sum)
		}
		if char == "D" {
			st, v1 = st.Pop()
			st = st.Push(v1)
			sum += v1 * 2
			st = st.Push(v1 * 2)
		}
		fmt.Println(sum)

	}
	return sum
}

// @lc code=end
