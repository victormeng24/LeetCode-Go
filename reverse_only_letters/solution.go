package reverse_only_letters

/**
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。

 

示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-only-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func reverseOnlyLetters(S string) string {
	arr := []byte(S)
	i, j := 0, len(arr) - 1
	for i < j {
		if !isLetter(arr[i]) {
			i++
			continue
		}
		if !isLetter(arr[j]) {
			j--
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}

func isLetter(x byte) bool {
	return (x >= 97 && x <= 122) || (x >= 65 && x <= 90)
}