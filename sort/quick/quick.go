package quick

func QuickSort(array []int) []int {
	quickSort(&array, 0, len(array)-1)

	return array
}

func quickSort(array *[]int, left, right int) {
	index := partition(array, left, right)

	if left < index-1 {
		quickSort(array, left, index-1)
	}

	if index < right {
		quickSort(array, index, right)
	}
}

func partition(array *[]int, left, right int) int {
	pivot := (*array)[(left+right)/2]

	for left <= right {
		for (*array)[left] < pivot {
			left++
		}

		for (*array)[right] > pivot {
			right--
		}

		if left <= right {
			(*array)[left], (*array)[right] = (*array)[right], (*array)[left]

			left++
			right--
		}
	}

	return left
}
