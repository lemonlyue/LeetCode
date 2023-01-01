package _53

/**
 * 剑指 Offer 53 - II. 0～n-1中缺失的数字
 * https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/description/
 */
func missingNumber(nums []int) int {
	res := len(nums)
	for index, item := range nums {
		if index != item {
			if index < item {
				return index
			} else {
				return item
			}
		}
	}
	return res
}
