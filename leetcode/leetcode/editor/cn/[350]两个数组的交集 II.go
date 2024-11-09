//给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现
//次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//
//
// 示例 2:
//
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000
//
//
//
//
// 进阶：
//
//
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
//
//
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 1060 👎 0

package main

// leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	// 基础版

		m := make(map[int]int)
		for _, v := range nums1 {
			m[v]++
		}
		for _, v := range nums2 {
			if m[v] > 0 {
				result = append(result, v)
				m[v]--
			}
		}

	// 进阶一：假设已经排好序
	/*slices.Sort(nums1)
	slices.Sort(nums2)

	var indexA, indexB = 0, 0
	for indexA < len(nums1) && indexB < len(nums2) {
		if nums1[indexA] > nums2[indexB] {
			indexB++
		} else if nums1[indexA] < nums2[indexB] {
			indexA++
		} else {
			result = append(result, nums1[indexA])
			indexA++
			indexB++
		}
	}
	*/
	// 进阶二：假设nums1的长度比nums2小
	/*// 把nums1（较短的那个数组）给哈希，然后nums2来判断是否存在
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}*/

	return result
}

//leetcode submit region end(Prohibit modification and deletion)
