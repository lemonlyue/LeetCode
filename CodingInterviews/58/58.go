package _58

/**
 * 左旋转字符串
 * https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/description/
 */
func reverseLeftWords(s string, n int) string {
	temp := s[0:n]
	s = s[n:]
	s = s + temp
	return s
}
