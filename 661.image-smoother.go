/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 */

// @lc code=start
import (
	"fmt"
)

func imageSmoother(M [][]int) [][]int {
	height := len(M)
	width := len(M[0])
	if height == 1 && width == 1 {
		return M
	}
	S := make([][]int, height)
	for i := 0; i < height; i++ {
		S[i] = make([]int, width)
	}
	//handle one dimension array
	if height == 1 {
		for j := 0; j < width; j++ {
			if j == 0 {
				S[0][j] = (M[0][j] + M[0][j+1]) / 2
			} else if j == width-1 {
				S[0][j] = (M[0][j] + M[0][j-1]) / 2
			} else {
				S[0][j] = (M[0][j-1] + M[0][j] + M[0][j+1]) / 3
			}
		}
	} else if width == 1 {
		for i := 0; i < height; i++ {
			if i == 0 {
				S[i][0] = (M[i][0] + M[i+1][0]) / 2
			} else if i == height-1 {
				S[i][0] = (M[i][0] + M[i-1][0]) / 2
			} else {
				S[i][0] = (M[i-1][0] + M[i][0] + M[i+1][0]) / 3
			}
		}
		//2X2 array and other larger arrays
	} else {
		for i := 0; i < height; i++ {
			column := []int{0, 0, 0}
			for j := 0; j < width; j++ {
				//corner
				if i == 0 && j == 0 {
					S[i][j] = (M[i][j] + M[i][j+1] + M[i+1][j] + M[i+1][j+1]) / 4
				} else if i == 0 && j == width-1 {
					S[i][j] = (M[i][j] + M[i][j-1] + M[i+1][j] + M[i+1][j-1]) / 4
				} else if i == height-1 && j == 0 {
					S[i][j] = (M[i-1][j] + M[i-1][j+1] + M[i][j] + M[i][j+1]) / 4
				} else if i == height-1 && j == width-1 {
					S[i][j] = (M[i-1][j] + M[i-1][j-1] + M[i][j] + M[i][j-1]) / 4
					//edge
				} else if i == 0 {
					S[i][j] = (M[i][j-1] + M[i][j] + M[i][j+1] + M[i+1][j-1] + M[i+1][j] + M[i+1][j+1]) / 6
				} else if i == height-1 {
					S[i][j] = (M[i][j-1] + M[i][j] + M[i][j+1] + M[i-1][j-1] + M[i-1][j] + M[i-1][j+1]) / 6
				} else if j == 0 {
					S[i][j] = (M[i-1][j] + M[i][j] + M[i+1][j] + M[i-1][j+1] + M[i][j+1] + M[i+1][j+1]) / 6
				} else if j == width-1 {
					S[i][j] = (M[i-1][j] + M[i][j] + M[i+1][j] + M[i-1][j-1] + M[i][j-1] + M[i+1][j-1]) / 6
				} else {
					//inside block
					//only compute the sum array for the first inside column, and reuse the stored computation result
					if j == 1 {
						for k := j - 1; k <= j+1; k++ {
							for l := i - 1; l <= i+1; l++ {
								column[k-j+1] += M[l][k]
							}
						}
					} else {
						tmpSum := 0
						for l := i - 1; l <= i+1; l++ {
							tmpSum += M[l][j+1]
						}
						column[0] = column[1]
						column[1] = column[2]
						column[2] = tmpSum
					}
					sum := 0
					for m := 0; m < 3; m++ {
						sum += column[m]
					}
					S[i][j] = sum / 9
					fmt.Print("Smooth tile row ", i, " column ", j, " is : ", S[i][j], "\n")
				}
			}
		}
	}
	return S
}

// @lc code=end
