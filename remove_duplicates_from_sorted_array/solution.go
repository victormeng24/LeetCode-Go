package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	len := len(nums)
	if len <= 1 {
		return len
	}
	res := 1
	for i := 1; i < len; i++ {
		if nums[i] > nums[i-1] {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
