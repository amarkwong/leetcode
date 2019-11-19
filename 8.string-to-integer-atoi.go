import "fmt"

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	return atoi(format(trimLeftWhiteSpaces(str)))
}

func trimLeftWhiteSpaces(str string) string {
	s := str
	for _, char := range str {
		if char == ' ' {
			s = s[1:]
		} else {
			return string(s)
		}
	}
	return string(s)
}
func atoi(str string) int {
	s := str
	if str[0] == '-' || str[0] == '+' {
		s = s[1:]
	}
	n := 0
	for i := 0; i < len(s); i++ {
		if n > 21478364 {
			if str[0] == '-' {
				return -2147483648
			} else {
				return 2147483647
			}
		} else if n == 21478364 {
			if str[0] == '-' {
				if s[i]-'0' >= 8 {
					return -2147483648
				} else {
					n = n*10 + int(s[i]-'0')
				}
			} else {
				if s[i]-'0' >= 7 {
					return 2147483647
				} else {
					n = n*10 + int(s[i]-'0')
				}
			}
		} else {
			n = n*10 + int(s[i]-'0')
		}
	}
	if str[0] == '-' {
		n = -n
	}
	return n
}

func format(str string) string {
	s := ""
	n := 0
	fmt.Println(len(s))
	if len(str) == 0 {
		//empty string
		fmt.Println("empty", s)
		return "0"
	} else if str[0] != '-' && str[0] != '+' && (str[0] < '0' || str[0] > '9') {
		//not valid format
		fmt.Println("invalid", s)
		return "0"
	} else {
		if str[0] == '-' {
			s = "-"
			fmt.Println("negative", s)
		} else if str[0] == '+' {
			s = "+"
			fmt.Println("positive", s)
		} else if str[0] >= '0' && str[0] <= '9' {
			s = "+"
			s += string(str[0])
			fmt.Println("positive number", s)
		} else {
			s = "0"
			return s
		}

		for i := 1; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				s += string(str[i])
				fmt.Println("loading numbers", s)
			} else {
				if len(s) == 0 {
					if str[i] == '-' || str[i] == '+' {
						s = string(str[i])
					} else {
						return "0"
					}
				} else {
					return s
				}
			}
		}
	}
	fmt.Println(s)
	return s
}

// @lc code=end
