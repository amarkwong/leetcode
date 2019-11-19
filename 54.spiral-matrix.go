/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	directions=[...]string{"right","down","left","up"}
	height,width:=len(matrix),len(matrix[0])
	if height==0||width==0 {
		return nil
	}
	res:=make([]int,0)
	leftright,updown:=1,0
	res=append(res,matrix[0][0])
	dim=height*width
	for i:=0;i<dim;i++ {
		if leftright==1 {
			for j=0;j<width;j++{
				if 
				res=append(res,matrix[i][j])
			}
		} else if leftright==0&&{
			for k=0;k<height;k--

		}
	}
}
// @lc code=end

