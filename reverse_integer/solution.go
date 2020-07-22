package reverse_integer

import "math"

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func reverse(x int) int {
	var tmp = int64(x)
	var y int64
	for tmp != 0 {
		y = y * 10 + tmp % 10
		tmp /= 10
		if y > math.MaxInt32 || y < math.MinInt32 {
			return 0
		}
	}
	return int(y)
}
