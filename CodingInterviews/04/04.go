package _4

/**
 * 二维数组中的查找
 * https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/description/
 */
func findNumberIn2DArray(matrix [][]int, target int) bool {
	res := false
	x, y := 0, 0
	xLength := len(matrix)

	if xLength <= 0 {
		return false
	}
	yLength := len(matrix[0])
	for true {
		if x >= xLength {
			break
		}
		if y >= yLength {
			x++
			y = 0
			continue
		}
		if matrix[x][y] < target {
			y++
			continue
		}
		if matrix[x][y] > target {
			x++
			y = 0
			continue
		}
		if matrix[x][y] == target {
			res = true
			break
		}
	}
	return res
}
