package merge_two_sorted_lists

import "testing"
import . "LeetCode-Go/common"

func TestSolution(t *testing.T)  {
	head := mergeTwoLists(BuildList([]int{1, 2, 4}), BuildList([]int{1, 3, 4}))
	PrintList(head)
}
