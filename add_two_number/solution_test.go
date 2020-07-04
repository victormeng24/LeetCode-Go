package add_two_number

import "testing"
import . "leetcode/common"

func TestSolution(t *testing.T) {
	result := addTwoNumbers(BuildList([]int{2, 4, 3}), BuildList([]int{5, 6, 4}))
	PrintList(result)
}
