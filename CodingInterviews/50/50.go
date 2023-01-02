package _50

/**
 * 剑指 Offer 50. 第一个只出现一次的字符
 * https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/description/
 */
func firstUniqChar(s string) byte {
	var keys []rune
	for _, item := range s {
		keys = append(keys, item)
	}
	data := make(map[rune]int)
	for _, item := range s {
		if data[item] != 0 {
			data[item]++
		} else {
			data[item] = 1
		}
	}

	for _, key := range keys {
		if data[key] == 1 {
			return byte(key)
		}
	}

	return ' '
}
