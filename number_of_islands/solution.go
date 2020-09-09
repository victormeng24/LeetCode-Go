package number_of_islands

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

 

示例 1:

输入:
[
['1','1','1','1','0'],
['1','1','0','1','0'],
['1','1','0','0','0'],
['0','0','0','0','0']
]
输出: 1
示例 2:

输入:
[
['1','1','0','0','0'],
['1','1','0','0','0'],
['0','0','1','0','0'],
['0','0','0','1','1']
]
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(&grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func dfs(grid *[][]byte, i int, j int)  {
	if i < 0 || i >= len(*grid) {
		return
	}
	if j < 0 || j >= len((*grid)[i]) {
		return
	}
	if (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	dfs(grid, i + 1, j)
	dfs(grid, i - 1, j)
	dfs(grid, i, j + 1)
	dfs(grid, i, j - 1)
}
