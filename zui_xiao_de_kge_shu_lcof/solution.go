package zui_xiao_de_kge_shu_lcof

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

 

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
 

限制：

0 <= k <= arr.length <= 10000
0 <= arr[i] <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	heap := arr[:k:k]
	for i := (k - 1) / 2; i >= 0; i-- {
		heapify(heap, i, k-1)
	}

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]

			heapify(heap, 0, k-1)
		}
	}
	return heap
}

func heapify(data []int, root int, end int) {
	for {
		child := root*2 + 1
		if child > end {
			return
		}
		for ; child < end && data[child] < data[child+1]; child++ {

		}
		if data[root] > data[child] {
			return
		}
		data[root], data[child] = data[child], data[root]
		root = child
	}
}
