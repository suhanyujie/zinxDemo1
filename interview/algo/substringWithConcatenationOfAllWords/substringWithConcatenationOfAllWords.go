package substringWithConcatenationOfAllWords

/**
题目名称：30. 串联所有单词的子串
题目地址：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/

## 题目描述
给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。

### 示例 1

```
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
```

### 示例 2

```
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
```

## 思路分析
本题采用滑动窗口+哈希表的解法实现，虽然不是最佳解法，但这种解法相对简单。

主要思路，先对特殊情况进行判断：
* 长度为 0 的 words，直接返回空切片。
* 其他情况，描述如下：

已知 words 中的单词长度都一样，我们可以先计算单词的总长度（单词个数 * 一个单词长度）记为 `allWordLen`

然后在 s 上，逐个取 allWordLen 个长度的子字符串，有了“子字符串”，我们接下来需要逐个匹配 words 中的单词是否都在其中，由于题目要求：单词完全匹配，中间不能有其他字符。
所以匹配过的字符不能再用于比较。我们先把 words 中的单词都存在一个 map 中（记为 map1），为了防止同一个单词出现多次的情况，我们对 map 的单词计数，因为类型为：`map[string]int`。

在每次比较“子字符串”时，会有一个新的 map 存储此次单词的匹配情况（记为 map2），从“子字符串”上截取一个单词长度，存入 map2 并计数，如果这个单词也在 map1 中，那么继续遍历，下一次比较的 offset 是 oneLen（即一个单词长度）。
如果不在，则跳过这个“子字符串”，继续取下一次的“子字符串”来比较。直到顺序取完所有的“子字符串”。

## 总结
解题时如果用的上 map 结构，是一种很好的辅助工具，因为 map 的存取复杂度是常数级，效率很高。

### 复杂度分析
* 循环遍历获取“子字符串”，循环长度依 s 的长度（记为 m）而定；第二层的单词匹配，根据 words 的长度（记为 n）而定。所以时间复杂度是 O(m*n)
* 空间复杂度：map 的存储空间由 words 的长度而定，因此为 O(n)

## ref
* 参考题解：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-6/

*/

// findSubstring
// s = "barfoothefoobarman", words = ["foo","bar"]
// 参考题解：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-6/
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	wordNum := len(words)
	if wordNum < 1 {
		return res
	}
	oneLen := len(words[0])
	allWordLen := len(words) * oneLen
	// map1 存储所有的 words
	map1 := make(map[string]int, len(words))
	for _, word := range words {
		map1[word] += 1
	}
	// 可以利用滑动窗口比较的次数
	cmpNum := len(s) - allWordLen + 1
	for i := 0; i < cmpNum; i++ {
		cmpStr := s[i : allWordLen+i]
		tmpNum := 0
		hasWordMap := make(map[string]int, 0)
		// 比较单词是否是否都能匹配上
		for j := 0; j < len(cmpStr)-oneLen+1; j += oneLen {
			tmpOneWord := cmpStr[j : oneLen+j]
			if _, ok := map1[tmpOneWord]; ok {
				hasWordMap[tmpOneWord] += 1
				if hasWordMap[tmpOneWord] > map1[tmpOneWord] {
					break
				}
			} else {
				break
			}
			tmpNum += 1
		}
		if tmpNum == wordNum {
			res = append(res, i)
		}
	}

	return res
}
