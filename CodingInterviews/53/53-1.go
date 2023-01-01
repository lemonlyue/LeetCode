package _53

/**
 * 在排序数组中查找数字 I
 * https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/description/
 */
func search(nums []int, target int) int {
	var num int

	// 暴力遍历
	for _, item := range nums {
		// 大于则跳出
		if item > target {
			return num
		}
		if item == target {
			num++
		}
	}
	return num
}
