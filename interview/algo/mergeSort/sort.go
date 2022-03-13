package mergeSort

/**
## 归并排序
令我萌生实现一下归并排序的来源是看了丁奇的“MySQL 实战 45 讲”的第 16 讲中的数据库排序问题，
在 MySQL 的实现中，当一条语句有排序需求时，如果无法用到覆盖索引的排序或者内存排序时，就可能会用到基于临时文件的归并排序。

## ref
* https://blog.csdn.net/qq_19381989/article/details/114387450

*/

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	div := len(nums) / 2
	left := mergeSort(nums[0:div])
	right := mergeSort(nums[div:])
	return mergeTwo(left, right)
}

// 合并两组
func mergeTwo(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	lLen, rLen := len(left), len(right)
	i, j := 0, 0
	for i < lLen && j < rLen {
		if left[i] > right[j] {
			result = append(result, right[j])
			j += 1
		} else {
			result = append(result, left[i])
			i += 1
		}
	}
	// 不满足 `i < lLen && j < rLen` 后，需要确定 left，right 是否还有剩余元素
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
