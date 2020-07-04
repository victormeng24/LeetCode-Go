package two_sum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if res, bol := m[target-v]; bol == true {
			return []int{res, i}
		}
		m[v] = i
	}
	return []int{-1, -1}
}
