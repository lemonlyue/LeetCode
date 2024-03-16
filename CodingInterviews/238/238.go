package _238

/**
238. 除自身以外数组的乘积
https://leetcode.cn/problems/product-of-array-except-self/description/
*/
// 除法
func productExceptSelf1(nums []int) []int {
	var res []int
	// 存储0的数量
	zeroNum := 0
	total := 1
	for _, num := range nums {
		if num == 0 {
			zeroNum++
			continue
		}
		total *= num
	}
	for _, num := range nums {
		// 判断0的数量是否大于1，大于1则都是0
		if zeroNum > 1 {
			res = append(res, 0)
			continue
		}
		// 判断0的数量是否为1
		if zeroNum == 1 {
			if num != 0 {
				res = append(res, 0)
			} else {
				res = append(res, total)
			}
			continue
		}
		// 否则按除法计算
		res = append(res, total/num)
	}
	return res
}

// 前后缀积
func productExceptSelf2(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for index, _ := range res {
		res[index] = 1
	}
	beforeSum := 1
	afterSum := 1
	i := 0
	j := length - 1
	for _, _ = range nums {
		res[i] *= beforeSum
		res[j] *= afterSum
		beforeSum *= nums[i]
		afterSum *= nums[j]

		i++
		j--
	}
	return res
}
