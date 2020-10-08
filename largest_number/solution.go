package largest_number

import (
	"sort"
	"strconv"
	"strings"
)

/**
给定一组非负整数 nums，重新排列它们每位数字的顺序使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

 

示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"
示例 3：

输入：nums = [1]
输出："1"
示例 4：

输入：nums = [10]
输出："10"
 

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func largestNumber(nums []int) string {
	numsStringSlice := make([]string, len(nums))
	for i, num := range nums {
		numsStringSlice[i] = strconv.Itoa(num)
	}
	sort.Slice(numsStringSlice, func(i, j int) bool {
		return numsStringSlice[j] + numsStringSlice[i] < numsStringSlice[i] + numsStringSlice[j]
	})
	ret := strings.Join(numsStringSlice, "")
	if ret[0] == '0' {
		return "0"
	}
	return ret
}