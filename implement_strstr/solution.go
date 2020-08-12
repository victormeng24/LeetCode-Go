package implement_strstr

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	return kmp(haystack, needle)
}

func kmp(s string, p string) int {
	next := make([]int, len(p))
	buildNext(next, p)
	j := 0
	for i := 0; i < len(s); i++ {
		for j > 0 && s[i] != p[j] {
			j = next[j-1] + 1
		}
		if s[i] == p[j] {
			j++
		}
		if j == len(p) {
			return i - j + 1
		}
	}
	return -1
}

func buildNext(next []int, s string) {
	j := -1
	next[0] = -1
	for i := 1; i < len(s); i++ {
		for j != -1 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
}
