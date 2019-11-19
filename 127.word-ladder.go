/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 */

//check whether the start string can be tranform to end string
func oneStep(start string, end string) int {
	if start == end {
		return 0
	}
	if len(start) != len(end) {
		return -1
	}
	res := 0
	for i := 0; i < len(start); i++ {
		if start[i] != end[i] {
			res++
		}
		if res > 1 {
			return -1
		}
	}
	return res
}

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//findout whether the endWord in the wordList
	//if not, return 0
	ms := make(map[string]int)
	me := make(map[string]int)
	for idx, str := range wordList {
		ms[str] = oneStep(beginWord, str)
		//take the ms with value 1 to generate next ms
		//if nothing with value 1 then return 0
		me[str] = oneStep(endWord, str)
		//check whether the me has value 0, if not return 0
		//if has value 0 but doesn't have value 1 return 0
		//if both exist select those with 1,
		//running the set with wordList
	}

	queue = make([]string, 0)

}

// @lc code=end
