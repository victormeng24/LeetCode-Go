package subsets

/**
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func subsets(nums []int) [][]int {
	var res [][]int
	var item []int
	helper(&item, &res, 0, nums)
	return res
}

func helper(tmp *[]int, res *[][]int, idx int, nums []int) {
	slice := make([]int, len(*tmp), len(*tmp))
	copy(slice, *tmp)
	*res = append(*res, slice)
	for i := idx; i < len(nums); i++ {
		*tmp = append(*tmp, nums[i])
		helper(tmp, res, i+1, nums)
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
