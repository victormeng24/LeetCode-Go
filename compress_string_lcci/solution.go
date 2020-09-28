package compress_string_lcci

import (
	"fmt"
	"strings"
)

/**
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:

 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：

字符串长度在[0, 50000]范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/compress-string-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func compressString(S string) string {
	var ret strings.Builder
	if len(S) <= 1 {
		return S
	}
	ret.WriteByte(S[0])
	cnt := 1
	for i := 1; i < len(S); i++ {
		if S[i] == S[i -1] {
			cnt++
		} else {
			ret.WriteString(fmt.Sprintf("%d", cnt))
			ret.WriteByte(S[i])
			cnt = 1
		}
	}
	ret.WriteString(fmt.Sprintf("%d", cnt))
	if ret.Len() > len(S) {
		return S
	}
	return ret.String()
}