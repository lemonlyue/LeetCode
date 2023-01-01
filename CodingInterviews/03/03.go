package _3

/**
 * 数组中重复的数字
 * https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/description/
 */
func findRepeatNumber(nums []int) int {
	data := make(map[int]int, len(nums))
	for _, item := range nums {
		if data[item] != 0 {
			return item
		} else {
			data[item] = 1
		}
	}
	return 0
}
