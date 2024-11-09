//给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
//
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
//
//
//
//
// 示例 1：
//
//
//输入：s = "abcdefg", k = 2
//输出："bacdfeg"
//
//
// 示例 2：
//
//
//输入：s = "abcd", k = 2
//输出："bacd"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由小写英文组成
// 1 <= k <= 10⁴
//
//
// Related Topics 双指针 字符串 👍 619 👎 0

package main

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	arr := strings.Split(s, "")
	length := len(arr)
	// 步长直接是2k
	for i := 0; i < length-1; i += 2 * k {
		// 设定左右边界
		left, right := i, i+k
		// 防止数组右边界越界
		if right >= length {
			right = length
		}
		// 注意middle取值
		for j := left; j < (left+right)/2; j++ {
			arr[j], arr[right - j - 1 + left] = arr[right - j - 1 + left], arr[j]
		}
	}

	return strings.Join(arr, "")

}

//leetcode submit region end(Prohibit modification and deletion)
