/*
 * @lc app=leetcode id=537 lang=golang
 *
 * [537] Complex Number Multiplication
 */

// @lc code=start
import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	aSplit := strings.Split(strings.Trim(a, "i"), "+")
	bSplit := strings.Split(strings.Trim(b, "i"), "+")
	var aReal, aImage, bReal, bImage int
	aReal, _ = strconv.Atoi(aSplit[0])
	aImage, _ = strconv.Atoi(aSplit[1])
	bReal, _ = strconv.Atoi(bSplit[0])
	bImage, _ = strconv.Atoi(bSplit[1])
	return strconv.Itoa((aReal*bReal - aImage*bImage)) + "+" + strconv.Itoa((aReal*bImage + aImage*bReal)) + "i"
}

// @lc code=end
