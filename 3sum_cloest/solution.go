package _sum_cloest

import (
	"math"
	"sort"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

 

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 

提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ret := -1
	minDiff := math.MaxInt32
	for i := 0; i < len(nums)-1; i++ {
		subTarget := target - nums[i]
		j, k := i+1, len(nums)-1
		for j < k {
			diff := nums[j] + nums[k] - subTarget
			if abs(diff) < minDiff {
				minDiff = abs(diff)
				ret = nums[i] + nums[j] + nums[k]
			}
			if diff > 0 {
				k--
			} else if diff < 0 {
				j++
			} else {
				return target
			}
		}

	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
