package maximum_subarray

/**

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
type Status struct {
	LSum int
	RSum int
	ISum int
	MSum int
}

func divideAndConquer(nums []int, l int, r int) Status {
	if l == r {
		return Status{
			LSum: nums[l],
			RSum: nums[l],
			ISum: nums[l],
			MSum: nums[l],
		}
	}
	mid := (l + r) >> 1
	lNode := divideAndConquer(nums, l, mid)
	rNode := divideAndConquer(nums, mid + 1, r)
	return Status{
		LSum: max(lNode.LSum, lNode.ISum + rNode.LSum),
		RSum: max(rNode.RSum, lNode.RSum + rNode.ISum),
		ISum: lNode.ISum + rNode.ISum,
		MSum: max(max(lNode.MSum, rNode.MSum), lNode.RSum + rNode.LSum),
	}
}
func maxSubArray(nums []int) int {
	return divideAndConquer(nums, 0, len(nums) - 1).MSum
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
