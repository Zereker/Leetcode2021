package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	if (len(nums1) + len(nums2)) == 1 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}

	locMed1 := (len(nums1) + len(nums2) - 1) / 2 // 中位数判定的第一个位置
	locMde2 := (len(nums1) + len(nums2) ) / 2 // 中位数判定的第二个位置
	loc, locA, locB, med1, med2 := 0, 0, 0, 0, 0 // loc：全局下标，locA:a游标，med1:中位数判定的第一个数

	for loc <= locMde2 { // 当总体的位置小于确定中位数的位置则循环
		if locA >= len(nums1) || len(nums1) == 0{
			nums1 = nil
		}

		if locB >= len(nums2) || len(nums2) == 0{
			nums2 = nil
		}

		if nums1 != nil && nums2 != nil {
			if nums1[locA] <= nums2[locB] {
				if loc == locMed1 {
					med1 = nums1[locA]
				}

				if loc == locMde2 {
					med2 = nums1[locA]
				}

				locA += 1
				loc += 1
				continue
			} else {
				if loc == locMed1 {
					med1 = nums2[locB]
				}

				if loc == locMde2 {
					med2 = nums2[locB]
				}

				locB += 1
				loc += 1
				continue
			}
		}

		if nums1 == nil && nums2 != nil{
			if loc == locMed1 {
				med1 = nums2[locB]
				locB += 1
				loc += 1
				continue
			}

			if loc == locMde2 {
				med2 = nums2[locB]
				locB += 1
				loc += 1
				continue
			}
		}

		if nums2 == nil && nums1 != nil{
			if loc == locMed1 {
				med1 = nums1[locA]
				locA += 1
				loc += 1
				continue
			}

			if loc == locMde2 {
				med2 = nums1[locA]
				locA += 1
				loc += 1
				continue
			}
		}

	}

	return (float64(med1) + float64(med2)) / 2.0

}

func main()  {
	r := findMedianSortedArrays([]int{3, 4}, []int{})
	fmt.Println(r)
}
