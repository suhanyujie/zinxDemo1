package nextPermutation

/**
题目名称：31. 下一个排列
题目地址：https://leetcode-cn.com/problems/next-permutation/

## 题目描述
题目来自[力扣](https://leetcode-cn.com/problems/next-permutation)，描述如下：

整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

* 例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。

整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

* 例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
* 类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
* 而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。

给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。

### 实例 1

```
输入：nums = [1,2,3]
输出：[1,3,2]
```
## 思路分析
根据[精选题解](https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-powcai/)
所谓的求解相邻的下一个排列可以视为（用相同数组中的数字）组合的相邻大数字。

例如：{1,2,3,4} 视为 `1234`，它的下一个大数值是 `1243`，而此时 {1,2,4,3} 恰好是题目中要求的相邻排列。
通过这种转换，我们就容易理解何为相邻的下一个排列。接下来，就要判断怎么操作（置换）才能得到期望的结果。主要步骤如下：

* 从左往右，找到最大的 i，使 nums[i] < nums[i+1]。其实就是**最靠近右侧**的相邻的升序对，以 {1,2,6,5,4,3} 为例，就是要找到其中的 (2,6) 这个升序对。
* 接着，从左往右，找到另一个**最靠近右侧**的索引 j，使其满足 nums[j] > nums[i]。也就是上方示例中的 3 值。
* 将 nums[i] 和 nums[j] 两值进行交换。上方示例，交换后，得到 {1,3,6,5,4,2}
* 然后，我们需要将 (i,end] 区间内的数值做升序排列（从左往右升序）（反转），得到：{1,3,2,4,5,6}，而这就是我们期望的结果。

## AC 代码

## 总结
在做完这题后，主要思路还是先弄明白相邻的下一组排列的规律，然后按照规律进行元素置换。值得一提的是，go 中可以通过赋值运算符实现元素的交换，加上双指针，可以实现切片的 reverse 操作。

### 复杂度分析
* 时间复杂度：先寻找最佳的相邻升序对，然后对一部分数值进行发转操作，没有多余的循环，可视为 O(n)
* 空间复杂度：所需的是确定个数的变量用于存储下标，空间复杂度是常数级 O(1)

## ref
* https://leetcode-cn.com/problems/next-permutation/

*/

func nextPermutation(nums []int) {
	length := len(nums)
	index := length - 1
	// 从后向前，查找升序对（这样查找，是最靠近右侧的）
	for ; index >= 0; index -= 1 {
		if nums[index-1] < nums[index] {
			break
		}
	}
	i := index - 1
	j := 0
	if i >= 0 {
		for j = length - 1; j > 0; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// reverse (i,end]
	right := length - 1
	for left := i + 1; left <= right; {
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
}
