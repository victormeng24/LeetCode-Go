package reverse_integer

import "math"

func reverse(x int) int {
	var tmp = int64(x)
	var y int64
	for tmp != 0 {
		y = y * 10 + tmp % 10
		tmp /= 10
		if y > math.MaxInt32 || y < math.MinInt32 {
			return 0
		}
	}
	return int(y)
}
