package valid_perfect_square

/**
给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。

说明：不要使用任何内置的库函数，如  sqrt。

示例 1：

输入：16
输出：True
示例 2：

输入：14
输出：False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-perfect-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	begin, end := 1, num/2
	for begin <= end {
		mid := begin + (end-begin)>>1
		tmp := mid * mid
		if tmp > num {
			end = mid - 1
		} else if tmp < num {
			begin = mid + 1
		} else {
			return true
		}
	}
	return false
}
