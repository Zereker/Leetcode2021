package main

/*
	https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	const max = 10 ^ 6

	len1 := len(nums1)
	len2 := len(nums2)
	total := len1 + len2

	m := 0
	n := 0
	i := 0

	var ileft float64
	var iright float64

	for m < len1 || n < len2 {
		num := 0

		n1 := max
		n2 := max

		if m < len1 {
			n1 = nums1[m]
		}

		if n < len2 {
			n2 = nums2[n]
		}

		if n1 < n2 {
			m += 1
			num = n1
		} else {
			n += 1
			num = n2
		}

		if total%2 == 0 {
			if i == total/2-1 {
				ileft = float64(num)
			}

			if i == total/2 {
				iright = float64(num)
			}

			if i > total/2 {
				break
			}

		} else {
			if i == total/2 {
				return float64(num)
			}
		}

		i++
	}

	return (ileft + iright) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)

	if total%2 == 0 {
		left := find(nums1, 0, nums2, 0, total/2)
		right := find(nums1, 0, nums2, 0, total/2+1)

		return float64(left+right) / 2.0
	}

	return float64(find(nums1, 0, nums2, 0, total/2+1))
}

func find(nums1 []int, i int, nums2 []int, j int, k int) int {
	if len(nums1)-i > len(nums2)-j {
		return find(nums2, j, nums1, i, k)
	}

	if len(nums1) == i {
		return nums2[j+k-1]
	}

	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	si, sj := min(i+k/2, len(nums1)), j+k-k/2
	if nums1[si-1] > nums2[sj-1] {
		return find(nums1, i, nums2, sj, k-(sj-j))
	} else {
		return find(nums1, si, nums2, j, k-(si-i))
	}
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}

	return i2
}
