package add_binary

import (
	"strconv"
)

/**
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。

 

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
 

提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func addBinary(a string, b string) string {
	var res = ""
	i, j := len(a)-1, len(b)-1
	var carry int
	for i >= 0 && j >= 0 {
		k := int(a[i]-'0') + int(b[j]-'0') + carry
		carry = k / 2
		res = strconv.Itoa(k%2) + res
		i--
		j--
	}
	for i >= 0 {
		k := int(a[i]-'0') + carry
		carry = k / 2
		res = strconv.Itoa(k%2) + res
		i--
	}
	for j >= 0 {
		k := int(b[j]-'0') + carry
		carry = k / 2
		res = strconv.Itoa(k%2) + res
		j--
	}
	if carry > 0 {
		res = "1" + res

	}
	return res
}
