package sort

func bubble(num []int) {
	for i := 0; i < len(num); i++ {
		for j := 1; j < len(num)-i; j++ {
			if num[j] < num[j-1] {
				num[j], num[j-1] = num[j-1], num[j]
			}
		}
	}
}

func insert(num []int) {
	for i := 1; i < len(num); i++ {
		j := i - 1
		value := num[i]

		for j >= 0 && num[j] > value {
			num[j+1] = num[j]
			j--
		}

		num[j+1] = value
	}
}

func selectSort(num []int) {
	for i := 0; i < len(num); i++ {
		maxIndex := 0
		for j := 1; j < len(num)-i; j++ {
			if num[j] > num[maxIndex] {
				maxIndex = j
			}
		}

		num[len(num)-i-1], num[maxIndex] = num[maxIndex], num[len(num)-i-1]
	}
}
