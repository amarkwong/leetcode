/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
func mapDigitToChar(digit int , k int) string{
	if digit<=7{
		return string('a'+ (digit -2)*3 + k )
	}else{
		return string('a'+ (digit -2)*3 + k + 1 )
	}
}
func letterCombinations(digits string) []string {
	if len(digits)==1 {
		digit:=int(digits[0]-'2')
		if digit==7||digit==9{
			
		}
	}

}
// @lc code=end

