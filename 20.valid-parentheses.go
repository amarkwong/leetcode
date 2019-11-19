import "fmt"

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
type stack []byte

func (s stack) Push(v byte) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, byte) {
	// FIXME: What do we do if the stack is empty, though?
	l := len(s)
	if l == 0 {
		return nil, '0'
	}
	return s[:l-1], s[l-1]
}

func Pair(b byte) byte {
	if b == '(' {
		return ')'
	} else if b == '[' {
		return ']'
	} else if b == '{' {
		return '}'
	} else {
		fmt.Print("invalid input")
		return byte(1)
	}
}

func isValid(s string) bool {
	st := make(stack, 0)
	if len(s)%2 == 1 {
		return false
	}
	for i, j := 0, byte(0); i < len(s); i++ {
		if len(st) == 0 {
			if s[i] == '(' || s[i] == '[' || s[i] == '{' {
				st = st.Push(s[i])
				j = Pair(s[i])
			} else {
				return false
			}
		} else {
			if j == byte(1) {
				return false
			} else if s[i] == '(' || s[i] == '[' || s[i] == '{' {
				st = st.Push(s[i])
				j = Pair(s[i])
			} else if s[i] == j {
				st, _ = st.Pop()
				if len(st) != 0 {
					j = Pair(st[len(st)-1])
				} else {
					j = byte(0)
				}
			} else {
				return false
			}
		}
	}
	if len(st) == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end
