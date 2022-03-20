package divideTwoIntegers

import (
	"math"
	"strconv"
)

/**
题目名称：29. 两数相除
题目地址：https://leetcode-cn.com/problems/divide-two-integers/

## 题目描述
题目来自[力扣](https://leetcode-cn.com/problems/divide-two-integers/)。描述如下：

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

### 示例 1

```
输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
```

### 示例 2：

```
输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2
```

## 思路分析
在一开始做这道题的时候，我感觉无从下手：
* 不用乘法、除法和 mod 运算符，那我还能用什么？
* 我也不知道如何处理除法结果溢出的情况

总之，我觉得大脑一片空白，我很气馁，但我想还是应该多看看题解，做题做多了也许就能看出一个题目考的是什么知识点。相信也有很多其他朋友也面临类似问题，很高兴能分享我的这些无知，
也很高兴能从无知到有知，如果你也有一些独特的想法，欢迎你在评论区写下你的感受。

回到题目本身。通过官方题解，我了解到，可以使用减法，看到这里，不用看完，我大概知道如何做了，实际的运算可以理解为：dividend 有多少个 divisor，我们通过减法，推断可以减去多少个 divisor，
这样，得到的次数，应该就是“商”了。但问题可不会就这么简单，我们需要考虑到另外两个问题：

* 1.被除数、除数的符号不一致如何处理？
* 2.何时出现结果溢出？

问题1，可以通过具体判断来解决。如果一正一负，我们可以使用 math.Abs() 让负数变为正数，然后返回结果前，让其变为负数。
问题2：根据题目提示除数、被除数的范围是 [−2^31,  2^31 − 1]，可知如果被除数为 −2^31，除数为 -1，则得到的结果是 2^31。
它无法用 int 类型表示，此时就溢出了。

如果只是使用这种简单的暴力枚举解法，放在力扣上执行，会发生超时情况，为了能顺利通过解题，只能寻找效率更高的解法。

主要的思路是先让 divisor 翻倍

* 如果被除数还是大，则让结果（商）翻倍，于此同时，被除数也需要减去双倍的除数，以进行下一次的运算。
* 如果翻倍后做减法 dividend 比双倍的 divisor 小，则结果为 1（递归的临界点）。

## 总结
也许是之前遇到的题过于片面，这才导致这次的第一次遇见执行超时情况。突然觉得算法的魅力在于比普通解法具有更高的效率，这才是刷题人应该追求的算法，而非使用暴力解法解题。
当然暴力解法也是解题的基础，甚至有时候还是关键点。

### 复杂度分析
* 只要不爆栈，随着 div 函数递归的进行，虽然除数不变，但是被除数会指数级减少，以 64 为例，变化过程形如 64->62->58->50->34，因此可以推测时间复杂度接近 O(log n)
* 空间复杂度，使用的变量数量是常数个，所以为 O(1)

*/

// divide 朴素解法
// 参考 https://leetcode-cn.com/problems/divide-two-integers/solution/po-su-de-xiang-fa-mei-you-wei-yun-suan-mei-you-yi-/
func divide(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	// 异常时的结果 2^31
	errCount := 2<<(31-1) - 1
	if divisor == -1 {
		// 特殊处理一下溢出情况
		if dividend <= math.MinInt32 {
			return errCount
		}
		return -dividend
	}
	if dividend == 0 {
		return 0
	}
	isNeg := false
	dividendF64 := float64(dividend)
	divisorF64 := float64(divisor)
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		isNeg = true
	}
	dividendF64 = math.Abs(dividendF64)
	divisorF64 = math.Abs(divisorF64)
	resF64 := div(dividendF64, divisorF64)
	resStr := strconv.FormatFloat(resF64, 'f', 0, 64)
	resInt64, _ := strconv.ParseInt(resStr, 10, 64)
	res := int(resInt64)
	if isNeg {
		return -res
	}
	if res > math.MaxInt32 {
		return errCount
	}

	return res
}

// 调用此方法时，a, b 都是大于 0 的。
func div(a, b float64) float64 {
	if a < b {
		return 0
	}
	count := float64(1)
	tmpB := b
	for a >= (tmpB + tmpB) {
		count = count + count
		tmpB = tmpB + tmpB
	}

	return count + div(a-tmpB, b)
}

// divide1 暴力破解
// Notice: 暴力破解的方式，在力扣执行会发生超时
func divide1(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}
	// 异常时的结果 2^31
	errCount := 2<<(31-1) - 1
	if divisor == -1 {
		// 特殊处理一下溢出情况
		if dividend <= math.MinInt32 {
			return errCount
		}
		return -dividend
	}
	if dividend == 0 {
		return 0
	}
	count := 0
	isNeg := false
	dividendF64 := float64(dividend)
	divisorF64 := float64(divisor)
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		dividendF64 = math.Abs(dividendF64)
		divisorF64 = math.Abs(divisorF64)
		isNeg = true
	}
	tmpDividend := dividendF64
	for tmpDividend >= divisorF64 {
		tmpDividend = tmpDividend - divisorF64
		count += 1
	}
	if isNeg {
		count = -count
	}
	if count > math.MaxInt32 {
		return errCount
	}

	return count
}
