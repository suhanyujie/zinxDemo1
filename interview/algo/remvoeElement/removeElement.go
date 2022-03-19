package remvoeElement

/**
题目名称：27. 移除元素
题目地址：https://leetcode-cn.com/problems/remove-element/

## 题目描述
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

### 示例 1

```
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
```

### 示例 2

```
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。
```

## 思路分析
这题与第 26 题很相似，一个是移除重复的，而这是移除指定的元素。个人感觉，这一题比上一题简单很多。
使用双指针，left 指针从索引 0 开始，right 指针从最后一位（`len(nums)-1`）开始。遍历 nums[left]

* `nums[left]` 与目标值 val 相等，如果 `nums[right] != val`，则将其赋值给 `nums[left]`，若 `nums[right] == val`，则 left 不变，right 左移。
* `nums[left]` 与目标值 val 不相等，则 left 右移，循环继续。

并将 right 下标左移。需要注意的是临界点，：

* 以 `[1,2,3]，target: 1` 为例，循环 1 次结束后，left 和 right 下标都指向值 2。此时 `nums[left] != val`，无需移动，`left += 1` 继续下一次循环。此后，`left > right`, 不满足条件，循环退出。
right 指向的索引为 1，即有效的数组长度为 2。
* 如果示例为 `[1,1,3]，target: 1`。循环 1 次结束后，left 和 right 下标都指向索引为 1 的元素 —— 1，left 最终值为 1，right 值指向索引为 0 的元素。即有效的数组长度为 1。

## 总结
主要的方向就是将尾部的有效元素移到前面，保证前面的元素都是有效的。此外，需要注意循环一次后，根据具体情况进行 left 右移，或者 right 左移。

无需额外的容器存储值，空间复杂度 O(1)；循环的次数根据 nums 的长度而定，因此，时间复杂度为 O(n)。

*/

// removeElement 移除元素
// 双指针解法
// [1,12,24,14,5,19,14,23]   14
func removeElement(nums []int, val int) int {
	length := len(nums)
	left, right := 0, length-1
	for left = 0; left <= right; {
		if nums[left] != val {
			left += 1
		} else {
			if nums[right] != val {
				nums[left] = nums[right]
				left += 1
			}
			right -= 1
		}
	}

	return right + 1
}
