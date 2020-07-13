package remove_duplicates_from_sorted_array

import "testing"

func TestSolution(t *testing.T)  {
	input := []int{0,0,1,1,1,2,2,3,3,4}
	res := removeDuplicates(input)
	t.Log(res)
}
