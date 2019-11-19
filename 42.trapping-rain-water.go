/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start

func floorSub(height []int) []int {
	for i := 0; i < len(height); i++ {
		if height[i] != 0 {
			height[i]--
		}
	}
	return height
}

func countDent(height []int, head int, tail int) (int, int, int) {
	// fmt.Print("Recounting Dent \n")
	// fmt.Print("Head: ", head, " Tail: ", tail, "\n")
	newHead, newTail, count := 0, 0, 0
	for i, j, leftBarrier, rightBarrier := head, tail-1, 0, 0; i < tail; i, j = i+1, j-1 {
		if height[i] == 0 {
			count++
		}
		if rightBarrier == 0 && height[j] != 0 {
			newTail = j + 1
			rightBarrier = 1
			// fmt.Print(" NEWTAIL: ", newTail, "\n")
		}
		if leftBarrier == 0 && height[i] != 0 {
			newHead = i
			leftBarrier = 1
		}
	}
	// fmt.Print("newHead: ", newHead, " newTail: ", newTail, "\n")
	// fmt.Print("Raw count: ", count, "\n")
	// fmt.Print("count: ", count-(newHead-head)-(tail-newTail), "\n")
	return count - (newHead - head) - (tail - newTail), newHead, newTail
}

func trap(height []int) int {
	trap, count := 0, 0
	if len(height) < 3 {
		return 0
	} else {
		for head, tail := 0, len(height); head < tail-2; height = floorSub(height) {
			count, head, tail = countDent(height, head, tail)
			trap += count
		}
		return trap
	}
}

// @lc code=end
