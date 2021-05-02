package main

/*
	https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/
/**
 * @Description:参考了题解，采用二分法找分割线
 */
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	//让num1做为短的数组
	if m > n {
		nums1, nums2 = nums2, nums1
	}
	totalLen := (m + n + 1) / 2
	//寻找分割线
	left, right := 0, len(nums1)
	for left < right {
		//左边的数要满足 nums1[i-1]<nums2[j]
		i := left + (right-left+1)/2
		j := totalLen - i

		//说明左边找的太大了，区间变为[left,i-1]
		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			//还可以继续往后找,区间变为[i,right]
			left = i
		}
	}

	i := left
	j := totalLen - i
	nums1LeftMax, nums2LeftMax := -1>>32, -1>>32
	if i != 0 {
		nums1LeftMax = nums1[i-1]
	}
	if j != 0 {
		nums2LeftMax = nums2[j-1]
	}
	nums1RightMin, nums2RightMin := 1<<31, 1<<31

	if i != len(nums1) {
		nums1RightMin = nums1[i]
	}
	if j != len(nums2) {
		nums2RightMin = nums2[j]
	}

	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2
	} else {
		return float64(max(nums1LeftMax, nums2LeftMax))
	}

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
