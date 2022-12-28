package _5

/**
 * 替换空格
 * https://leetcode.cn/problems/ti-huan-kong-ge-lcof/description/
 */
func replaceSpace(s string) string {
	ans := ""
	for _, item := range s {
		switch item {
		case ' ':
			ans = ans + "%20"
		default:
			ans = ans + string(item)
		}
	}
	return ans
}
