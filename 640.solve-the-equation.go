/*
 * @lc app=leetcode id=640 lang=golang
 *
 * [640] Solve the Equation
 */

// @lc code=start
import (
	"strings"
)
func solveEquation(equation string) string {


}
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	equation := "1+3x-6x+18-9+10-6x+5+12-19=34x-12+13"
	E := strings.Split(equation, "=")
	fmt.Println(E[0], E[1])
	cl := strings.Split(E[0], "-")
	fmt.Println(cl)
	c := 0
	ex := 0
	for i := 0; i < len(cl); i++ {
		cm := strings.Split(cl[i], "+")
		fmt.Println("without +", cm)
		for j := 0; j < len(cm); j++ {
			if strings.HasSuffix(cm[j], "x") {
				if j == 0 {
					if i == 0 && !(strings.HasPrefix(cl[0], "-")) {
						tmp, _ := strconv.Atoi(strings.TrimRight(cm[j], "x"))
						ex += tmp
					} else {
						tmp, _ := strconv.Atoi(strings.TrimRight("-"+cm[j], "x"))
						ex += tmp
					}
				} else {
					tmp, _ := strconv.Atoi(strings.TrimRight(cm[j], "x"))
					ex += tmp
				}
			} else {
				if j == 0 {
					if i == 0 && !(strings.HasPrefix(cl[0], "-")) {
						tmpc, _ := strconv.Atoi(cm[j])
						fmt.Println("first time positive c :",c)	
						c += tmpc
					} else {
						fmt.Println("first time negative c :",c)	
						tmpc, _ := strconv.Atoi("-" + cm[j])
						c += tmpc
					}
				} else {
					fmt.Println("not first time c :",c)	
					tmpc, _ := strconv.Atoi(cm[j])
					c += tmpc
				}
				fmt.Println("c :",c)	
			}
		}
	}
	fmt.Print("constant: ", c, " coefficient: ", ex)
}


// @lc code=end

