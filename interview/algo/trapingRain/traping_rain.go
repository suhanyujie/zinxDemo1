package trapingRain

/**
>* 题目名称：42. 接雨水
>* 题目地址：https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/

## 题目描述
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

### 示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

### 示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

## 分析
讲真，做这种题我一般第一感觉是去暴力破解。但是看来官方题解后，发现解法还是有很多的，而且效率还很好。比较好的解法是双指针解法。但直接上双指针会难以理解。

所以，还是先理解一下动态规划解法吧。

先明确一下变量概念，从左往右，用 leftMax[i] 表示索引为 0~i 之间的最大元素值；类似的，从右往左，用 rightMax[i] 表示 i~len(height)-1 之间的元素最大值。

通过两次循环遍历，得到每个下标所对应的 leftMax 和 rightMax，并以相同的下标为 key，存储到数组中备用，记为 leftMaxArr 和 rightMaxArr。

然后从左往右再遍历一次，分别比较 leftMaxArr 和 rightMaxArr 在下标 i 处的值，取最小的那个，减去当前 i 处的 height[i] 值，就是下标 i 处所能存储的单位水量。将其累加，即可得到总的蓄水量。

## 总结
动态规划理解起来虽然也不容易，但按照思路多走几次，逻辑上还是可理解的。双指针解法是在此基础上更进一步，只用两个变量存储 leftMax 和 rightMax，空间复杂度上提升为常数级，值得学习！

## ref
* https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/

*/

// trap 双指针解法
func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	i, j := 0, len(height)-1
	leftMax, rightMax := height[i], height[j]
	total := 0
	for i < j {
		if height[i] > leftMax {
			leftMax = height[i]
		}
		if height[j] > rightMax {
			rightMax = height[j]
		}
		if height[i] < height[j] {
			total += leftMax - height[i]
			i += 1
		} else {
			total += rightMax - height[j]
			j -= 1
		}
	}

	return total
}

// trapDp 动态规划解法
func trapDp(height []int) int {
	if len(height) < 1 {
		return 0
	}
	length := len(height)
	// 两个数组存储两边的 side max
	leftMaxArr, rightMaxArr := make([]int, length), make([]int, length)
	leftMax, rightMax := height[0], height[length-1]
	for i, val := range height {
		if val > leftMax {
			leftMax = val
		}
		leftMaxArr[i] = leftMax
	}
	for j := length - 1; j >= 0; j-- {
		if height[j] > rightMax {
			rightMax = height[j]
		}
		rightMaxArr[j] = rightMax
	}
	total := 0
	// 从左往右，leftMax[i]
	for i := 0; i < length; i++ {
		total += Min(leftMaxArr[i], rightMaxArr[i]) - height[i]
	}

	return total
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
