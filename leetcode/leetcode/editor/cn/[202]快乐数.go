//编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」 定义为：
//
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
//
//
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
//
//
//
// 示例 1：
//
//
//输入：n = 19
//输出：true
//解释：
//1² + 9² = 82
//8² + 2² = 68
//6² + 8² = 100
//1² + 0² + 0² = 1
//
//
// 示例 2：
//
//
//输入：n = 2
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
//
// Related Topics 哈希表 数学 双指针 👍 1615 👎 0

package main

import (
	"fmt"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	// 解析：可能无限循环，那就是因为平方和出现了重复，所以需要一个东西来保存每次平方和的结果
	// 首先得能找到int每个位置的数字，我考虑是把int转成一个可以遍历的数组
	allNumSet := make(map[int]void)  // 这里是利用map实现一个set，value是空的结构体
	for n != 1 {
		if _, ok := allNumSet[n]; ok {
			break
		}
		// 记录每次平方和的数字
		allNumSet[n] = member
		var result int
		// 这是通过strconv在数字和字符串之间转换来获取每个位置上的数字的
		/*str := strconv.Itoa(n)
		for _, ch := range str {
			num, _ := strconv.Atoi(string(ch))
			result += num * num
		}*/
		// 这是通过模和除来计算每个位置上的数字的
		for n > 0 {
			num := n % 10
			result += num * num
			n /= 10
		}
		// 迭代更新n的值
		n = result
		fmt.Println(n)
	}
	fmt.Println(allNumSet)
	return n == 1

}

type void struct{}

var member void
// leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(isHappy(2))
}
