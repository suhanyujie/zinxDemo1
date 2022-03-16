package removeDuplicatesFromSortedArr

/**
题目名称：26. 删除有序数组中的重复项
题目地址：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

## 题目描述
以下是从[ leetcode 网站](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)摘录的题目大意：

> 给你一个 升序排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。元素的相对顺序应该保持一致。
由于在某些语言中不能改变数组的长度，所以必须将结果放在数组 nums 的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
将最终结果插入 nums 的前 k 个位置后返回 k 。
不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。判题标准:

> 系统会用下面的代码来测试你的题解:

```
int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```

> 如果所有断言都通过，那么您的题解将被 通过。

### 示例

```
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
```

## 分析
说来也巧，在做这道题的时候，我是在排队做核酸检测的，看到双指针关键词的时候，我还没看题解，我当时已经在心里打好草稿了，大概能走通了。但是回到家，发现有些地方竟想不通。
原因是看到错误的代码，竟按照错误的思路一直走，难怪走不通。所以，其他题解还是少看，多看官方题解，如果看不懂，再看其他参考。

对于这题，官方题解描述的还是挺好的，一开始双指针我用的变量是 `i`, `j`，但后来觉得 `slow`, `fast` 变量名更形象。

* 对于长度小于 2 的输入，我们可以直接返回 `len(nums)`。
* 长度大于 2 时，第 `0` 个元素我们永远不用更改，通过分类套路可知：第 0, 1 个元素相同时，只需修改第 `1` 个元素的值；如果第 0, 1 个元素不同时，则也无需更改第 1 个元素。

我们可以将 slow 指针视为“历史元素”的边界，即在它之前的元素要么是不同的，要么是同值却要保留下来的，因此当 `nums[fast] != nums[fast-1]` 时，需要将 `nums[fast]` 往前挪，即
覆盖 `nums[slow]` 元素。

## 总结
发现双指针的解法在很多题目中都能用上，它的确是解题必备的方法之一。如果对其不熟练，尝试多做几题双指针题，也许就能找到感觉了。

## ref
* https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

*/

// removeDuplicates 双指针写解法
//   1, 1, 1, 3, 6
//   1, 2, 7, 8
//
func removeDuplicates(nums []int) int {
	lenNum := len(nums)
	if lenNum < 2 {
		return lenNum
	}
	// 第一个数可以永远不用修改
	slow, fast := 1, 1
	for ; fast < lenNum; fast += 1 {
		// i==0 时，不能和 j 位置交换
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow += 1
		}
	}

	return slow
}
