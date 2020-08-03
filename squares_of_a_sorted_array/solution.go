package squares_of_a_sorted_array

/**
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

 

示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]
 

提示：

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A 已按非递减顺序排序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func sortedSquares(A []int) []int {
	i := 0
	for i < len(A) && A[i] < 0 {
		i++
	}
	j := i - 1
	res := make([]int, 0, len(A))
	for i < len(A) && j >= 0 {
		if A[i]*A[i] > A[j]*A[j] {
			res = append(res, A[j] * A[j])
			j--
		} else {
			res = append(res, A[i] * A[i])
			i++
		}
	}
	for i < len(A) {
		res = append(res, A[i] * A[i])
		i++
	}
	for j >= 0 {
		res = append(res, A[j] * A[j])
		j--
	}
	return res
}
