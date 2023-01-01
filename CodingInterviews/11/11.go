package _11

/**
 * 剑指 Offer 11. 旋转数组的最小数字
 * https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/description/
 */
func minArray(numbers []int) int {
	length := len(numbers)
	for index, item := range numbers {
		if index == length-1 {
			break
		}
		if item > numbers[index+1] {
			return numbers[index+1]
		}
	}
	return numbers[0]
}
