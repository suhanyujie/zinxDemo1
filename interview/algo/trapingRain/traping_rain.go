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

## 总结


## ref
* https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode-solution-tuvc/

*/

// trap 采用双指针解法
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
